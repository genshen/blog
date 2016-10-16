/**
 * Created by 根深 on 2016/8/14.
 */
var CONFIG = {
    adminRouter: "/admin",
    apiPrefix: "/admin/api",
    adminAuthPath: "/admin/auth/signin",
    adminSignOutPath: "/admin/auth/signout",
    adminStaticPrefix: "/private"
};

function init() {
    Util.postData.config.authUrl = CONFIG.adminAuthPath;

    $.get(CONFIG.adminStaticPrefix + "/templates/index.html", function (data) {
        document.getElementById("template-container").outerHTML = data;

        registerVueRouter();
    });
}

function registerVueRouter() {
    var App = Vue.extend({
        template: '#app_template',
        data: function () {
            return {}
        }, ready: function () {
        },
        methods: {}
    });

    var Menu = Vue.extend({
        template: '#menu-container',
        methods: {},
        ready: function () {
            console.log("menu");
        }
    });

    var ArticleList = Vue.extend({
        template: '#article-list',
        methods: {},
        ready: function () {
            console.log("list");
        }
    });

    var ArticleEdit = Vue.extend({
        template: '#article-edit',
        data: function () {
            return {
                upload_config:{token:"",upload_path:"",domain:""},
                image_uploading_processing:false,
                images:[], //{src:"blobUrl",status:0,file:fileObject}  //status:0未上传,1正在上传,2上传完成
                markedStatus: false,
                article_title: "",
                article_content: ""
            }
        },
        methods: {
            getUploadToken:function(){
                $.ajax({url:CONFIG.apiPrefix+"/upload_token",context:this,success:function(data){
                        this.upload_config = data;
                },error:function(req,err){
                        $("body").snackbar({alive: 3000, content: "加载上传配置信息出错了"});
                }});
            },
            addUploadImage:function(){
                if(this.upload_config.token){ //check upload_token
                    $("#upload_image_input").trigger("click");
                }else{
                    $("body").snackbar({alive: 3000, content: "UploadToken无效"});
                }
            },
            onUploadImageSelected:function(){
                var files = $("#upload_image_input")[0].files;
                var base_length = this.images.length;
                for(var i=0;i<files.length;i++){
                    var src = window.URL.createObjectURL(files[i]);
                    this.images.push({src: src,status:0, file: files[i]});
                    this.uploadImageToServer(this.images[base_length+i]);
                }
            },
            uploadImageToServer:function(image){
               var data = new FormData();
                data.append("token",this.upload_config.token);
                data.append("file",image.file);
                $.ajax({
                    url: this.upload_config.upload_path,
                    type: 'POST',
                    data: data,
                    context:this,
                    cache: false,
                    processData: false,
                    contentType: false,
                    success:function(data){
                        try{
                            this.article_content += "![image]("+this.upload_config.domain+data.key+")\r\n";
                        }catch(e){
                            $("body").snackbar({alive: 3000, content: "上传出错了"});
                        }
                    },error:function(){
                        $("body").snackbar({alive: 3000, content: "上传出错了"});
                    }
                });

            },
            // deleteUploadImage:function(index){
            //     if (index < this.images.length) {
            //         this.images.splice(index, 1)
            //     }
            // },
            submit: function () {
                if (!this.article_title) {
                    $("body").snackbar({content: "标题不能为空", alive: 4000});
                    return;
                } else if (!this.article_content) {
                    $("body").snackbar({content: "内容不能为空", alive: 4000});
                    return;
                }
                var self = this;
                Util.postData.init(CONFIG.apiPrefix + "/post/add/", {  //todo category_id
                    category_id:"57cd6bfe20953120fc2beafc",title: this.article_title, content: this.article_content,
                    summary: marked(this.article_content).replace(/<.*?>/ig, "")
                }, null, function () {
                    $("body").snackbar({content: "文章发布成功", alive: 4000});
                    self.article_title = "";
                    self.article_content = ""
                });
            }
        },
        filters: {
            markdown: function (content) {
                if (this.markedStatus) {
                    return marked(content);
                } else {
                    return content;
                }
            }
        },
        ready: function () {
            if (!this.markedStatus) {
                var self = this;
                loadJS(["//cdnjs.cloudflare.com/ajax/libs/marked/0.3.2/marked.min.js",
                    "//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.6.0/highlight.min.js"], function () {
                    marked.setOptions({
                        highlight: function (code) {
                            return hljs.highlightAuto(code).value;
                        }
                    });
                    self.markedStatus = true;
                });
            }
            //get image upload token
            this.getUploadToken();
        }
    });

    var CategorySetting = Vue.extend({
        template: '#settings-category',
        data: function () {
            return {
                categories: [],
                new_category_name: "",
                new_category_slug: "",
                new_category_submitting: false,
                new_sub_category_type: 0,
                new_sub_category_name: "",
                new_sub_category_slug: "",
                new_sub_category_submitting: false
            }
        },
        methods: {
            addCategory: function () {
                if (!this.new_category_submitting) {
                    if (this.new_category_name == "" || this.new_category_slug == "") {
                        $("body").snackbar({alive: 3000, content: "请填写完相关项后再提交"});
                        return;
                    }
                    this.new_category_submitting = true;
                    var self = this;
                    Util.postData.init(CONFIG.apiPrefix + "/category/add", {
                            name: this.new_category_name,
                            slug: this.new_category_slug
                        },
                        null, function (data) {
                            self.categories.push({
                                id: data.Addition,
                                name: self.new_category_name,
                                slug: self.new_category_name,
                                sub_category: []
                            });
                            self.new_category_name = "";
                            self.new_category_name = "";
                            $("#category-edit-modal").modal("hide");
                            $("body").snackbar({alive: 3000, content: "分类添加成功"});
                        }, null, null, null, function () {
                            self.new_category_submitting = false;
                        });
                }
            }, addSubCategory: function () {
                // console.log(this.new_sub_category_type);
                if (this.new_sub_category_name == "" || this.new_sub_category_slug == "") {
                    $("body").snackbar({alive: 3000, content: "请填写完相关项后再提交"});
                    return;
                }
                var index = this.new_sub_category_type;
                var _id = this.categories[index].id;
                this.new_sub_category_submitting = true;
                var self = this;

                Util.postData.init(CONFIG.apiPrefix + "/sub_category/add", {
                        id: _id,
                        name: this.new_sub_category_name,
                        slug: this.new_sub_category_slug
                    },
                    null, function (data) {
                        self.categories[index].sub_category.push({
                            id: data.Addition,
                            name: self.new_sub_category_name,
                            slug: self.new_sub_category_slug,
                            posts_count: 0
                        });
                        self.new_sub_category_name = "";
                        self.new_sub_category_slug = "";
                        $("#sub-category-edit-modal").modal("hide");
                        $("body").snackbar({alive: 3000, content: "子分类添加成功"});
                    }, null, null, null, function () {
                        self.new_sub_category_submitting = false;
                    });
            }
        },
        ready: function () {
            var self = this;
            $.ajax({
                url: CONFIG.apiPrefix + "/categories",
                success: function (data) {
                    try {
                        if (data) {
                            data.forEach(function (e) {
                                self.categories.push(e);
                            });
                        }
                    } catch (e) {
                        $("body").snackbar({alive: 3000, content: "出了点错误,请重试"});
                    }
                }, error: function (err) {
                    $("body").snackbar({alive: 3000, content: "出了点错误,请重试"});
                }
            });
        }
    });


    var router = new VueRouter({root: "/admin", hashbang: false, history: true});
    router.map({
        '/': {
            name: 'menu',
            component: Menu
        },
        '/article/list': {
            name: "article_list",
            component: ArticleList
        }, '/article/edit': {
            name: "article_edit",
            component: ArticleEdit
        }, '/settings/category': {
            name: "settings_category",
            component: CategorySetting
        }
    });

    router.start(App, 'template');
    if (document.location.pathname == "/admin") {
        router.go({name: "menu"})
    }
}