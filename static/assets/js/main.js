/**
 * Created by 根深 on 2016/8/11.
 */
var settings = {};
var Config = {
    apiPrefix: "/at"
};
var NoAuthSnackBar = '<a data-dismiss="snackbar">Dismiss</a>' +
    '<div class="snackbar-text">你需要登录认证后才能添加评论.' +
    '(使用<a href="https://github.com" target="_blank">Github</a>账号登录)</div>';

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
            settings = data;

            (function () {

                var App = Vue.extend({
                    template: '#app_template',
                    data: function () {
                        return {
                            settings: settings,
                            title: "哈哈哈",
                            detail: {is_auth: true}
                        }
                    }, ready: function () {
                    },
                    methods: {
                        openGithub: function () {
                            window.open("https://github.com");
                            $('#auth_model').modal('hide');
                        }
                    }
                });

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
                    created: function () {
                        var self = this;
                        $.ajax({
                            url: Config.apiPrefix + "/category",
                            success: function (data) { //if it is not json?
                                try {
                                    data.forEach(function (e) {
                                        if (!e.cover) {
                                            e.cover = "/assets/img/brand.jpg";
                                        }
                                        self.lists.push(e);
                                    })
                                } catch (err) { //todo
                                    console.log(err);
                                }

                            }, error: function (err) { //todo
                                console.log(err);
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
                            article: "# Marked in browser\n\nRendered by **marked**.\n```c\n int main(){\r\n if(i<o){j++;\r\nreturn 0}}\n```",
                            comments: [
                                {
                                    user: {name: "he", url: "baidu.com", avatar: "/assets/img/avatar.jpg"},
                                    replies: [
                                        {
                                            user: {name: "he", url: "baidu.com", avatar: "/assets/img/avatar.jpg"},
                                            content: "Hello replies1",
                                            date: "3天前"
                                        },
                                        {
                                            user: {name: "he", url: "baidu.com", avatar: "/assets/img/avatar.jpg"},
                                            content: "Hello replies2",
                                            date: "3天前"
                                        }
                                    ],
                                    content: "Hello Comment",
                                    date: "3天前",
                                    show_reply_box: false,
                                    new_reply_content: ""
                                },
                                {
                                    user: {name: "he", url: "baidu.com", avatar: "/assets/img/avatar.jpg"},
                                    replies: [],
                                    content: "Hello Comment2",
                                    date: "3天前",
                                    show_reply_box: false,
                                    new_reply_content: ""
                                }
                            ]
                        }
                    },
                    computed: {},
                    filters: {
                        marked: function (value) {
                            return marked(value);
                        }
                    },
                    methods: {
                        checkAuth: function () {
                            if (this.is_auth) {
                                return
                            }
                            $('#auth_model').modal('show')
                        },
                        submitComment: function () {
                            if (!this.is_auth) {
                                $("body").snackbar({alive: 4000, content: NoAuthSnackBar});
                                return
                            }
                            console.log("here! ");
                        },
                        submitReply: function (commentIndex) {
                            console.log("submitReply! ");
                        },
                        cancelReply: function (commentIndex) {
                            this.comments[commentIndex].show_reply_box = false;
                        },
                        toggleReplyBox: function (commentIndex, replyIndex) { //-1 ->is comment //todo isAuth
                            var atOne;
                            if (replyIndex < 0) {
                                atOne = this.comments[commentIndex].user.name
                            } else {
                                atOne = this.comments[commentIndex].replies[replyIndex].user.name
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
                        console.log("detail");
                    }
                });


                var router = new VueRouter({hashbang: false, history: true});
                router.map({
                    '/': {
                        component: List
                    },
                    '/detail/:id': {
                        name:"detail",
                        component: Detail
                    }
                });
                router.start(App, 'template');
            })();
        });
    });
}
