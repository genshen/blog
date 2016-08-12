/**
 * Created by 根深 on 2016/8/11.
 */
var settings = {};
var NoAuthSnackBar = '<a data-dismiss="snackbar">Dismiss</a>' +
    '<div class="snackbar-text">你需要登录认证后才能添加评论.' +
    '(使用<a href="https://github.com" target="_blank">Github</a>账号登录)</div>';
function init() {
    marked.setOptions({
        highlight: function (code) {
            return hljs.highlightAuto(code).value;
        }
    });
    $("body").load("/static/t/index.html", function () {
        $.get("/static/t/setting.json", function (data) {
            settings = data;
            settings.show_content_header = true;

            (function () {
                Vue.component('list-grid', {
                    template: '#list-template',
                    props: {
                        settings: Object,
                        show: Boolean
                    },
                    data: function () {
                        return {
                            lists: [{
                                date: "3天前",
                                title: "Hello World",
                                summary: "Lorem ipsum dolor sit amet.Consectetur adipiscing elit.",
                                img: "/static/img/brand.jpg"
                            }]
                        };
                    },
                    computed: {},
                    methods: {},
                    ready: function () {

                    }
                });
                Vue.component('detail-grid', {
                    template: '#detail-template',
                    props: {
                        settings: Object,
                        show: Boolean,
                        is_auth: Boolean //sync
                    },
                    data: function () {
                        return {
                            article: "# Marked in browser\n\nRendered by **marked**.\n```c\n int main(){\r\n if(i<o){j++;\r\nreturn 0}}\n```",
                            comments: [
                                {
                                    user: {name: "he", url: "baidu.com", avatar: "/static/img/avatar.jpg"},
                                    replies:[
                                        {
                                            user: {name: "he", url: "baidu.com", avatar: "/static/img/avatar.jpg"},
                                            content: "Hello replies1",
                                            date: "3天前"
                                        },
                                        {
                                            user: {name: "he", url: "baidu.com", avatar: "/static/img/avatar.jpg"},
                                            content: "Hello replies2",
                                            date: "3天前"
                                        }
                                    ],
                                    content: "Hello Comment",
                                    date: "3天前",
                                    new_reply:"",
                                    show_reply_box:true
                                },
                                {
                                    user: {name: "he", url: "baidu.com", avatar: "/static/img/avatar.jpg"},
                                    replies:[

                                    ],
                                    content: "Hello Comment2",
                                    date: "3天前"
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
                        }
                    },
                    ready: function () {

                    }
                });
            })();

            new Vue({
                el: "html",
                data: {
                    settings: settings,
                    title: "哈哈哈",
                    list: {show: false},
                    detail: {show: true, is_auth: true}
                },
                ready: function () {
                },
                methods: {
                    openGithub: function () {
                        window.open("https://github.com");
                        $('#auth_model').modal('hide')
                    }
                }
            });
        });
    });
}
