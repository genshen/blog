/**
 * Created by 根深 on 2016/8/14.
 */
var CONFIG = {
    adminRouter: "/admin",
    apiPrefix: "/admin/api",
    adminAuthPath: "/admin/auth/signin",
    adminSignOutPath: "/admin/auth/signout"
};
function init() {
    Util.postData.config.authUrl = CONFIG.adminAuthPath;

    $.get("/private/templates/index.html", function (data) {
        document.getElementById("template-container").outerHTML = data;

        registerComponent();
        new Vue({
            el: 'html',
            data: {
                currentView: "menu-container"
            }, methods: {
                changeView: function (viewName) {
                    this.currentView = viewName;
                }
            }
        })
    });
}

function registerComponent() {
    Vue.component('menu-container', {
        template: '#menu-container',
        props: {},
        methods: {},
        ready: function () {
            console.log("menu");
        }
    });

    Vue.component('article-list', {
        template: '#article-list',
        props: {},
        methods: {},
        ready: function () {
            console.log("list");
        }
    });

    Vue.component('article-edit', {
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
                Util.postData.init(CONFIG.apiPrefix + "/post/add/", {
                    title: this.article_title, content: this.article_content,
                    summary: marked(this.article_content).replace(/<.*?>/ig, "")
                }, null, function () {
                    console.log("success");
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
}