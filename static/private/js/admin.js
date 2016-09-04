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
        props: {},
        methods: {},
        ready: function () {
            console.log("menu");
        }
    });

    var ArticleList = Vue.extend({
        template: '#article-list',
        props: {},
        methods: {},
        ready: function () {
            console.log("list");
        }
    });

    var ArticleEdit = Vue.extend({
        template: '#article-edit',
        props: {},
        data: function () {
            return {
                markedStatus: false,
                article_title: "",
                article_content: ""
            }
        },
        methods: {
            submit: function () {
                if (!this.article_title) {
                    $("body").snackbar({content: "标题不能为空", alive: 4000});
                    return;
                } else if (!this.article_content) {
                    $("body").snackbar({content: "内容不能为空", alive: 4000});
                    return;
                }
                var self = this;
                Util.postData.init(CONFIG.apiPrefix + "/post/add/", {
                    title: this.article_title, content: this.article_content,
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
        }
    });
    router.start(App, 'template');
    if (document.location.pathname == "/admin") {
        router.go({name: "menu"})
    }
}