/*util：
 * including:
 * 1.base64_decode
 * 2.jquery.cookie.js
 * 3.js load
 * */

//base64_decode
function base64_decode(encodedData) {
    if (typeof window !== 'undefined') {
        if (typeof window.atob !== 'undefined') {
            return decodeURIComponent(unescape(window.atob(encodedData)))
        }
    } else {
        return new Buffer(encodedData, 'base64').toString('utf-8')
    }

    var b64 = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';
    var o1, o2, o3, h1, h2, h3, h4, bits;
    var i = 0;
    var ac = 0;
    var dec = '';
    var tmpArr = [];

    if (!encodedData) {
        return encodedData
    }

    encodedData += '';

    do {
        // unpack four hexets into three octets using index points in b64
        h1 = b64.indexOf(encodedData.charAt(i++));
        h2 = b64.indexOf(encodedData.charAt(i++));
        h3 = b64.indexOf(encodedData.charAt(i++));
        h4 = b64.indexOf(encodedData.charAt(i++));

        bits = h1 << 18 | h2 << 12 | h3 << 6 | h4;

        o1 = bits >> 16 & 0xff;
        o2 = bits >> 8 & 0xff;
        o3 = bits & 0xff;

        if (h3 === 64) {
            tmpArr[ac++] = String.fromCharCode(o1)
        } else if (h4 === 64) {
            tmpArr[ac++] = String.fromCharCode(o1, o2)
        } else {
            tmpArr[ac++] = String.fromCharCode(o1, o2, o3)
        }
    } while (i < encodedData.length);

    dec = tmpArr.join('');
    return decodeURIComponent(escape(dec.replace(/\0+$/, '')))
}

//js-cookie
/*!
 * JavaScript Cookie v2.1.2
 * https://github.com/js-cookie/js-cookie
 *
 * Copyright 2006, 2015 Klaus Hartl & Fagner Brack
 * Released under the MIT license
 */
(function (factory) {
    if (typeof define === 'function' && define.amd) {
        define(factory);
    } else if (typeof exports === 'object') {
        module.exports = factory();
    } else {
        var OldCookies = window.Cookies;
        var api = window.Cookies = factory();
        api.noConflict = function () {
            window.Cookies = OldCookies;
            return api;
        };
    }
}(function () {
    function extend() {
        var i = 0;
        var result = {};
        for (; i < arguments.length; i++) {
            var attributes = arguments[i];
            for (var key in attributes) {
                result[key] = attributes[key];
            }
        }
        return result;
    }

    function init(converter) {
        function api(key, value, attributes) {
            var result;
            if (typeof document === 'undefined') {
                return;
            }
            // Write
            if (arguments.length > 1) {
                attributes = extend({
                    path: '/'
                }, api.defaults, attributes);

                if (typeof attributes.expires === 'number') {
                    var expires = new Date();
                    expires.setMilliseconds(expires.getMilliseconds() + attributes.expires * 864e+5);
                    attributes.expires = expires;
                }

                try {
                    result = JSON.stringify(value);
                    if (/^[\{\[]/.test(result)) {
                        value = result;
                    }
                } catch (e) {
                }

                if (!converter.write) {
                    value = encodeURIComponent(String(value))
                        .replace(/%(23|24|26|2B|3A|3C|3E|3D|2F|3F|40|5B|5D|5E|60|7B|7D|7C)/g, decodeURIComponent);
                } else {
                    value = converter.write(value, key);
                }

                key = encodeURIComponent(String(key));
                key = key.replace(/%(23|24|26|2B|5E|60|7C)/g, decodeURIComponent);
                key = key.replace(/[\(\)]/g, escape);

                return (document.cookie = [
                    key, '=', value,
                    attributes.expires && '; expires=' + attributes.expires.toUTCString(), // use expires attribute, max-age is not supported by IE
                    attributes.path && '; path=' + attributes.path,
                    attributes.domain && '; domain=' + attributes.domain,
                    attributes.secure ? '; secure' : ''
                ].join(''));
            }

            // Read

            if (!key) {
                result = {};
            }

            // To prevent the for loop in the first place assign an empty array
            // in case there are no cookies at all. Also prevents odd result when
            // calling "get()"
            var cookies = document.cookie ? document.cookie.split('; ') : [];
            var rdecode = /(%[0-9A-Z]{2})+/g;
            var i = 0;

            for (; i < cookies.length; i++) {
                var parts = cookies[i].split('=');
                var cookie = parts.slice(1).join('=');

                if (cookie.charAt(0) === '"') {
                    cookie = cookie.slice(1, -1);
                }

                try {
                    var name = parts[0].replace(rdecode, decodeURIComponent);
                    cookie = converter.read ?
                        converter.read(cookie, name) : converter(cookie, name) ||
                    cookie.replace(rdecode, decodeURIComponent);

                    if (this.json) {
                        try {
                            cookie = JSON.parse(cookie);
                        } catch (e) {
                        }
                    }

                    if (key === name) {
                        result = cookie;
                        break;
                    }

                    if (!key) {
                        result[name] = cookie;
                    }
                } catch (e) {
                }
            }

            return result;
        }

        api.set = api;
        api.get = function (key) {
            return api(key);
        };
        api.getJSON = function () {
            return api.apply({
                json: true
            }, [].slice.call(arguments));
        };
        api.defaults = {};

        api.remove = function (key, attributes) {
            api(key, '', extend(attributes, {
                expires: -1
            }));
        };

        api.withConverter = init;

        return api;
    }

    return init(function () {
    });
}));

var Util = {
    postData: {
        config: {
            authUrl: "" //todo
        },
        init: function (url, data, o, onPostSuccess, onUnAuth, onPostError, onError) {
            var options = $.extend({}, {snackBarAlive: 4000,multiError:true, showNext: false, authUrl: this.config.authUrl}, o);
            if (!onError) {
                onError = function () {
                    $("body").snackbar({
                        content: "出了点错误,请<a href='" + window.location.href + "'>刷新</a>重试",
                        alive: options.snackBarAlive
                    });
                }
            }
            if (!onPostError) {
                if(options.multiError){
                    onPostError = function (Errors) {
                        for (var key in Errors) {
                            var err = Errors[key].Errors;
                            if (err.length > 0) {
                                $("body").snackbar({content:err[0].Message, alive: options.snackBarAlive});
                                return;
                            }
                        }
                    }
                }else{
                    onPostError = function (error) {
                        $("body").snackbar({content: error, alive: options.snackBarAlive});
                    }
                }
            }
            if (!onUnAuth) {
                onUnAuth = function () {
                    var url;
                    if (options.showNext) {
                        url = options.authUrl + "?next=" + +document.location.pathname;
                    } else {
                        url = options.authUrl;
                    }
                    $("body").snackbar({content: "请<a href='" + url + "'>登录</a>后进行操作", alive: options.snackBarAlive});
                }
            }
            this.execute(url, data, onPostSuccess, onUnAuth, onPostError, onError);
        },
        execute: function (url, data, onPostSuccess, onUnAuth, onPostError, onError) {
            var xsrf;
            try { //cookie may be null or something else bad data
                xsrf = base64_decode(Cookies.get('_xsrf').split("|")[0]);
            } catch (err) {
                onError();
                return;
            }
            $.ajax({
                type: 'POST',
                url: url,
                data: $.extend({}, {_xsrf: xsrf}, data),
                success: function (data) {
                    try {
                        switch (data.Status) {
                            case 0:
                                onPostError(data.Error);
                                break;
                            case 1:
                                if (onPostSuccess) {
                                    onPostSuccess();
                                }
                                break;
                        }
                    } catch (err) {
                        onError();
                    }
                }, error: function (r, err) {
                    if (r.status == 401) {
                        onUnAuth();
                    } else {
                        onError();
                    }
                }
            });
        }
    },
    simpleParseError: {
        options: {},
        init: function (status, error, options) {
            this.options = $.extend({}, {
                snackAlive: 3000,
                errorCallback: function (message) {
                    $("body").snackbar({content: message, alive: this.snackTimeout})
                },
                onSuccess: null
            }, options);
            this.execute(status, error)
        },
        execute: function (status, error) {
            switch (status) {
                case 0:
                    this.options.errorCallback(error);
                    return false;
                case 1:
                    if (this.options.onSuccess != null) {
                        this.options.onSuccess();
                    }
                    return true;
            }
        }
    }
};

//js load
function loadJS(srcs, cb) {
    "use strict";
    var total = srcs.length;
    var hasLoadedCount = 0;

    function onLoadCallback() {
        hasLoadedCount++;
        if (total == hasLoadedCount && cb && typeof(cb) === "function") {
            cb();
        }
    }

    for (var index in srcs) {
        var src = srcs[index];
        var ref = document.getElementsByTagName("script")[0];
        var script = document.createElement("script");
        script.src = src;
        script.async = true;
        ref.parentNode.insertBefore(script, ref);
        script.onload = onLoadCallback;
    }
}
function loadCSS(src, cb) {
    "use strict";
    var head = document.getElementsByTagName('HEAD').item(0);
    var style = document.createElement('link');
    style.href = src;
    style.rel = 'stylesheet';
    style.type = 'text/css';
    head.appendChild(style);
    if (cb && typeof(cb) === "function") {
        style.onload = cb;
    }
    return style;
}