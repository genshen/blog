/**
 * Created by 根深 on 2016/8/11.
 */
var settings = {};
function init() {
    $("body").load("/static/t/index.html", function () {
        $.get("/static/t/setting.json", function (data) {
            settings = data;
            settings.show_content_header = true;

            (function () {
                Vue.component('list-grid', {
                    template: '#list-template',
                    props: {
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
                        lists: Array,
                        show: Boolean
                    },
                    computed: {},
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
                    // this.list.data.push({
                    //     date: "3天前",
                    //     title: "Hello World",
                    //     summary: "Lorem ipsum dolor sit amet.Consectetur adipiscing elit.",
                    //     img: "/static/img/brand.jpg"
                    // });
                }
            });
        });
    });
}
