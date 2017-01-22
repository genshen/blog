/**
 * Created by 根深 on 2016/8/11.
 */
var settings = {}; //referred to vue(read only)
var categories;
var Config = {
    apiPrefix: "/at",
    MaxCommentLength: 20
};

var NormalErrorSnackBar = '<a data-dismiss="snackbar">Dismiss</a>' +
    '<div class="snackbar-text">Oh,Snap! 出了点错误,请' +
    '<a href="javascript:location.reload() ">刷新</a>' + '后重试.</div>';

var NoAuthCommentSnackBar = '<a data-dismiss="snackbar">Dismiss</a>' +
    '<div class="snackbar-text">你需要' +
    '<a data-toggle="modal" data-target="#auth_model">登录认证</a>' + '后才能添加评论.</div>';

var NoAuthReplySnackBar = '<a data-dismiss="snackbar">Dismiss</a>' +
    '<div class="snackbar-text">你需要' +
    '<a data-toggle="modal" data-target="#auth_model">登录认证</a>' + '后才能添加回复.</div>';

function init() {
    marked.setOptions({
        highlight: function (code) {
            return hljs.highlightAuto(code).value;
        }
    });
    $.get("/assets/t/index.html", function (data) {
        $("div").remove();
        document.getElementById("template-container").outerHTML = data;

        $.get("/settings", function (data) {
            settings = data.settings;
            categories = data.categories;
            settings.is_auth = data.is_auth;
            if (data.user) {
                settings.user = data.user;
            }

            (function () {
                Vue.filter('formatTime', formatTime);

                var List = Vue.extend({
                    template: '#list-template',
                    data: function () {
                        return {
                            settings: settings,
                            lists: []
                        };
                    },
                    computed: {},
                    methods: {},
                    route: {
                        data: function(){
                            console.log("data!@");
                        }
                    },
                    created: function () {
                        console.log("ready!@");
                        var self = this;
                        $.ajax({
                            url: Config.apiPrefix + "/category", //todo different category
                            success: function (data) { //if it is not json?
                                try {
                                    data.forEach(function (e) {
                                        if (!e.cover) {
                                            e.cover = "/assets/img/brand.jpg";
                                        }
                                        self.lists.push(e);
                                    })
                                } catch (err) {
                                    $("body").snackbar({alive: 3000, content: NormalErrorSnackBar});
                                }
                            }, error: function (err) {
                                $("body").snackbar({alive: 3000, content: NormalErrorSnackBar});
                            }
                        });
                    }
                });

                var Detail = Vue.extend({
                    template: '#detail-template',
                    props: {
                        is_auth: Boolean //sync
                    },
                    data: function () {
                        return {
                            settings: settings,
                            article: {
                                id: "", title: "", content: "", summary: "", cover: "",
                                view_count: 0, comment_count: 0, reply_count: 0, created_at: "", updated_at: ""
                            },
                            comments: [],
                            comments_loaded_count: 0,
                            comment_load_status: -1, //0 for loading(init),1 for loading, 2 for loaded(has more),
                            // 3 for loaded(no more),4 for failed to load(init),5 for failed to load
                            comment_text: ""
                        }
                    },
                    computed: {
                        compiledMarkdown: function () {
                            return marked(this.article.content);
                        }
                    },
                    methods: {
                        setCommentLoad: function () {
                            if (document.getElementById("comment-flag").offsetTop < document.documentElement.clientHeight) {
                                this.loadComment();
                            } else {
                                var self = this;
                                $(window).scroll(function () {
                                    if (document.getElementById("comment-flag").offsetTop - document.body.scrollTop <
                                        document.documentElement.clientHeight) {
                                        $(window).unbind('scroll');
                                        self.loadComment();
                                    }
                                });
                            }
                        },
                        loadComment: function () {
                            if (this.comment_load_status == 0 || this.comment_load_status == 1) {
                                return;
                            }
                            var start = this.comments_loaded_count;
                            this.comment_load_status = start == 0 ? 0 : 1;
                            var self = this;
                            $.ajax({
                                url: Config.apiPrefix + "/comments/" + this.$route.params.id + "/" + start,
                                success: function (data) {
                                    try {
                                        data.forEach(function (e) {
                                            if (!self.containsComment(e.id)) {
                                                e.show_reply_box = false;
                                                e.new_reply_content = "";
                                                self.comments.unshift(e);
                                            }
                                        });
                                        self.comments_loaded_count += data.length;
                                        self.comment_load_status = data.length == Config.MaxCommentLength ? 2 : 3; //2(has more) or 3(not more)
                                    } catch (err) {
                                        self.comment_load_status = start == 0 ? 4 : 5;
                                    }
                                }, error: function (r, err) {
                                    self.comment_load_status = start == 0 ? 4 : 5; //4 or 5
                                }
                            });
                        },
                        containsComment: function (id) {
                            var h = this.comments.length - 1, l = 0;
                            while (l <= h) {
                                var m = Math.floor((h + l) / 2);
                                if (this.comments[m].id == id) {
                                    return true;
                                }
                                if (id > this.comments[m].id) {
                                    l = m + 1;
                                } else {
                                    h = m - 1;
                                }
                            }
                            return false;
                        },
                        submitComment: function () {
                            if (this.comment_text == "") {
                                $("body").snackbar({alive: 3000, content: "评论内容不能为空"});
                                return;
                            }
                            if (!this.settings.is_auth) {
                                $('#auth_model').modal('show');
                                return;
                            }

                            var self = this;
                            Util.postData.init(Config.apiPrefix + "/comment/add/", {
                                post_id: this.article.id, content: this.comment_text
                            }, null, function (data) {
                                self.comments.unshift({
                                    content: self.comment_text, create_at: (new Date()).getTime(),
                                    id: data.Addition, replies: [], status: 1, user: settings.user,
                                    show_reply_box: false, new_reply_content: ""
                                });
                                self.comment_text = "";
                                $("body").snackbar({content: "评论成功", alive: 3000});
                            }, function (error) {
                                $("body").snackbar({alive: 3000, content: NoAuthCommentSnackBar});
                            });
                        },
                        submitReply: function (commentIndex) {
                            var self = this.comments[commentIndex];
                            if (self.new_reply_content == "") {
                                $("body").snackbar({alive: 3000, content: "回复内容不能为空"});
                                return;
                            }
                            if (!this.settings.is_auth) {
                                $('#auth_model').modal('show');
                                return;
                            }

                            Util.postData.init(Config.apiPrefix + "/reply/add/", {
                                comment_id: self.id, content: self.new_reply_content
                            }, null, function (data) {
                                try {
                                    self.replies.push({
                                        content: self.new_reply_content, create_at: (new Date()).getTime(),
                                        id: data.Addition, status: 1, user: settings.user
                                    });
                                    self.new_reply_content = "";
                                    self.show_reply_box = false;
                                    $("body").snackbar({content: "回复成功", alive: 3000});
                                } catch (e) {
                                    $("body").snackbar({content: "出了点错误,请重试", alive: 3000});
                                }
                            }, function (error) {
                                $("body").snackbar({alive: 3000, content: NoAuthReplySnackBar});
                            });
                        },
                        cancelReply: function (commentIndex) {
                            this.comments[commentIndex].show_reply_box = false;
                        },
                        toggleReplyBox: function (commentIndex, replyIndex) { //-1 ->is comment
                            var atOne;
                            if (replyIndex < 0) {
                                atOne = this.comments[commentIndex].user.name;
                            } else {
                                atOne = this.comments[commentIndex].replies[replyIndex].user.name;
                            }
                            this.comments[commentIndex].show_reply_box = true;
                            this.comments[commentIndex].new_reply_content = "@" + atOne;
                            // console.log($("#reply-box-"+commentIndex));
                            setTimeout(function () {
                                var box = $("#reply-box-" + commentIndex);
                                box.trigger("change");
                                box.trigger("focus");
                            }, 200);
                        }
                    },
                    created: function () {
                        var self = this;
                        $.ajax({
                            url: Config.apiPrefix + "/detail/" + this.$route.params.id,
                            success: function (data) { //if it is not json?
                                try {
                                    if (!data.id) {
                                        $("body").snackbar({alive: 4000, content: "Oh,Snap! 查看的文章不存在"});
                                        return;
                                    }
                                    if (!data.cover) {
                                        data.cover = "/assets/img/brand.jpg";
                                    }
                                    self.article = data;
                                } catch (err) {
                                    $("body").snackbar({alive: 3000, content: NormalErrorSnackBar});
                                }
                            }, error: function (err) {
                                $("body").snackbar({alive: 3000, content: NormalErrorSnackBar});
                            }
                        });
                        setTimeout(this.setCommentLoad, 100);
                    }
                });

                const router = new VueRouter({mode:"history",
                    routes: [
                        {path:'/', name: 'home', component: List},
                        {path:'/category/:menu/:sub_menu', name: 'category', component: List},
                        {path:'/detail/:id', name: "detail", component: Detail}
                    ]
                });

               new Vue({router:router,
                   data: function () {
                       return {
                           settings: settings,
                           categories: categories,
                           title: "哈哈哈",
                           detail: {is_auth: true}
                       }
                   }, created: function () {
                       console.log(this.settings);
                   },
                   methods: {
                       openGithub: function () {
                           var url = this.settings.auth_sites.github.url + this.settings.auth_sites.github.client_id;
                           window.open(url, "", "location=no,status=no");
                           $('#auth_model').modal('hide');
                       }
                   }
               }).$mount('#app');
            })();
        });
    });
}

window.addEventListener('message', function (e) {
    if (e.origin == location.origin) {
        var data = e.data;
        if (data.status == 1) {
            settings.is_auth = true;
            settings.user = data;
            $("body").snackbar({alive: 3000, content: "登录认证成功"});
        }
    }
});

function formatTime(value) {
    if (typeof value != "number") {
        var v = Date.parse(value);
        if (isNaN(v)) {
            value = (new Date).getTime();
        } else {
            value = v;
        }
    }
    now = (new Date).getTime();
    if (now - value < 60 * 1000) {
        return "刚刚";
    }
    if (now - value < 60 * 60 * 1000) {
        var min = parseInt((now - value) / (60 * 1000));
        return min + "分钟前";
    }
    if (now - value < 24 * 60 * 60 * 1000) {
        var hour = parseInt((now - value) / (60 * 60 * 1000));
        return hour + "小时前";
    }
    if (now - value < 20 * 24 * 60 * 60 * 1000) {
        var day = parseInt((now - value) / (24 * 60 * 60 * 1000));
        return day + "天前";
    }
    var d = new Date(value);
    return d.getFullYear()+"-"+(d.getMonth()+1)+"-"+d.getDate();
}