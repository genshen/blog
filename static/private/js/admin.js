/**
 * Created by 根深 on 2016/8/14.
 */
function init() {
    $.get("/private/templates/index.html", function (data) {
        document.getElementById("template-container").outerHTML = data;
    });
}