/**
 * Created by 根深 on 2016/8/11.
 */
var settings = {};
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
                        lists: Array,
                        show: Boolean
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
                        lists: Array,
                        show: Boolean
                    },
                    data: function () {
                        return {article: "# Marked in browser\n\nRendered by **marked**.\n```c\n int main(){\r\n if(i<o){j++;\r\nreturn 0}}\n```"}
                    },
                    computed: {},
                    filters: {
                        marked: function(value){
                            return marked(value);
                        }
                    },
                    methods: {},
                    ready: function () {

                    }
                });
            })();

            new Vue({
                el: "html",
                data: {
                    settings: settings,
                    title: "哈哈哈",
                    list: {show: false, data: []},
                    detail: {show: true, data: []}
                },
                ready: function () {
                    this.list.data.push({
                        date: "3天前",
                        title: "Hello World",
                        summary: "Lorem ipsum dolor sit amet.Consectetur adipiscing elit.",
                        img: "/static/img/brand.jpg"
                    });
                }
            });
        });
    });
}
