$(function () {});

//js设置cookie 并且全站有效
function set_cookie(name, value, liveMinutes) {
  // if (liveMinutes == undefined || liveMinutes == null) {
  //     liveMinutes = 60 * 2;
  // }
  // if (typeof (liveMinutes) != 'number') {
  //     liveMinutes = 60 * 2;//默认120分钟
  // }
  var minutes = liveMinutes * 60 * 1000;
  var exp = new Date();
  exp.setTime(exp.getTime() + minutes + 8 * 3600 * 1000);
  //path=/表示全站有效，而不是当前页
  document.cookie = name + "=" + value + ";path=/;expires=" + exp.toUTCString();
}

// 获取cookies值
function get_cookie(c_name) {
  if (document.cookie.length > 0) {
    c_start = document.cookie.indexOf(c_name + "=");
    if (c_start != -1) {
      c_start = c_start + c_name.length + 1;
      c_end = document.cookie.indexOf(";", c_start);
      if (c_end == -1) c_end = document.cookie.length;
      return unescape(document.cookie.substring(c_start, c_end));
    }
  }
  return "";
}

function Remove_Program(pName, Table) {
  // 从cookies移除字段
  var SelectProgram = getCookie(Table); // 读取己选中的数据
  var arr = SelectProgram.split(",");
  SelectProgram = ""; // 清空整体数据
  for (
    var i = 0;
    i < arr.length;
    i++ // 数据重新组合
  ) {
    if (arr[i] != pName && arr[i] != "" && arr[i] != null) {
      SelectProgram += arr[i] + ","; // 重组数据
    }
  }
  if (SelectProgram[SelectProgram.length - 1] == ",") {
    SelectProgram = SelectProgram.substr(0, SelectProgram.length - 1);
  }
  set_cookie(Table, SelectProgram); // 写入
  //console.log('RemoveProgram:'+SelectProgram);
}

function Add_Program(pName, Table) {
  // 向cookies添加字段
  var SelectProgram = getCookie(Table); // 读取己选中的数据
  var arr = SelectProgram.split(",");
  // alert(Table+' Add前:'+SelectProgram);
  SelectProgram = ""; // 清空整体数据
  for (
    var i = 0;
    i < arr.length;
    i++ // 数据重新组合
  ) {
    if (arr[i] != pName && arr[i] != "" && arr[i] != null) {
      SelectProgram += arr[i] + ","; // 重组数据
    }
  }
  SelectProgram += pName;
  if (SelectProgram[SelectProgram.length - 1] == ",") {
    SelectProgram = SelectProgram.substr(0, SelectProgram.length - 1);
  }
  set_cookie(Table, SelectProgram); // 写入
  //console.log('AddProgram:'+SelectProgram);
}

//前端网站index nav导航栏滑动
function navSlide() {
  $(".nav li").each(function () {
    var flag = 0;
    var i = 0;
    if ((get_cookie("flag") == "") & (get_cookie("i") == "")) {
      set_cookie("flag", 0);
      set_cookie("i", 0);
    }
    $(".nav li").mouseover(function () {
      var index = $(this).index();
      $(".nav .slide_line").css("left", (index - 1) * 94 + "px");
      $(".nav li a").css("color", "#333333");
      $(this).find("a").css("color", "#00B3DA");
      $(".nav li").click(function () {
        $(".nav .slide_line").css("left", (index - 1) * 94 + "px");
        flag = (index - 1) * 94;
        i = $(this).index() - 1;
        $(this).find("a").css("color", "#00B3DA");
        set_cookie("flag", flag);
        set_cookie("i", i);
      });
      $(".nav li").mouseout(function () {
        $(".nav .slide_line").css("left", get_cookie("flag") + "px");
        $(".nav li a").css("color", "#333333");
        $(".nav li a").eq(get_cookie("i")).css("color", "#00B3DA");
      });
    });
  });
}

//footer高度首页和其他页面不一致
function footerHeight() {
  $("#footer").css("height", "340px");
  $(".footer_center").css("height", "340px");
  $(".contact_us_p1").css("padding-top", "100px");
  $(".contact_icon .qq, .contact_icon .weixin").css("padding-top", "170px");
  $(".company_name").css("padding-top", "100px");
  $(".company_profile_list").css("height", "180px");
}

//邮箱验证  可以为空
function checkEmail(email) {
  if (email != "" && email != null) {
    // var pattern=/^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$/;
    var pattern = /^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$/;
    if (pattern.test(email)) {
      return true;
    } else {
      return false;
    }
  } else return true;
}

//判断是否含有非法字符
function findSpecialChar(char) {
  var pattern =
    /^SELECT|select|DELETE|delete|UPDATE|update|INSERT|insert|from|into|WHERE|where|AND|and|DROP|drop|table|database|alert|values|union|into|outfile$/;
  if (pattern.test(char)) {
    return false;
  } else return true;
}

//判断是否是正整数 可以为空
function checkNum(num) {
  if (num != "" && num != null) {
    var pattern = /^\+?[0-9][0-9]*$/;
    if (pattern.test(num)) {
      return true;
    } else {
      return false;
    }
  } else return true;
}

//判断邮政编码 可以为空
function checkZipcode(num) {
  if (num != "" && num !== null) {
    var pattern = /^[a-zA-Z0-9 ]{3,12}$/;
    if (pattern.test(num)) {
      return true;
    } else {
      return false;
    }
  } else return true;
}

//判断电话  可以为空
function checkPhone(num) {
  if (num != "" && num !== null) {
    // var isTel=/^([0-9]{3,4}-)?[0-9]{7,8}$/;    //固定电话
    // var isPhone=/^((\+?86)|(\(\+86\)))?(13[012356789][0-9]{8}|15[012356789][0-9]{8}|18[02356789][0-9]{8}|147[0-9]{8}|1349[0-9]{7})$/;  //手机
    var isPhone = /^1[345678]\d{9}$/; //手机
    if (isPhone.test(num)) {
      return true;
    } else {
      return false;
    }
  } else return true;
}

//正则表达 0到100 以内整数
function checkIntRange(num) {
  if ($.trim(num) != "" && $.trim(num) != null) {
    var pattern = /^([1-9][0-9]{0,1}|0|100)$/;
    if (pattern.test(num)) {
      return true;
    } else {
      return false;
    }
  } else return false;
}

//正则 年月日  yyyy-mm-dd 2008-11-01   也可以是中文  星期日 星期一
function checkYTD(str) {
  if ($.trim(str) != "" && $.trim(str) != null) {
    var pattern =
      /^((((19|20)\d{2})-(0?[13-9]|1[012])-(0?[1-9]|[12]\d|30))|(((19|20)\d{2})-(0?[13578]|1[02])-31)|(((19|20)\d{2})-0?2-(0?[1-9]|1\d|2[0-8]))|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))-0?2-29))$/;
    var type = /^[\u4e00-\u9fa5]+$/;
    if (pattern.test(str) || type.test(str)) {
      return 1;
    } else {
      return -2;
    }
  } else return -1;
}

//正则 小时:分钟 - 小时:分钟  时间段梵文  不能为空
function checkHoursMinutes(str) {
  if ($.trim(str) != "" && $.trim(str) != null) {
    // var pattern = /^([0-1]{1}\d|2[0-3]):([0-5]\d)$/;
    var pattern =
      /^([0-1]{1}\d|2[0-3]):([0-5]\d)\s-\s([0-1]{1}\d|2[0-3]):([0-5]\d)$/;
    if (pattern.test(str)) {
      return 1;
    } else {
      return -2;
    }
  } else return -1;
}

//获取textarea标签中的换行符和空格 html标签存进数据库
function getFormatCode(value) {
  return value
    .replace(/\r\n/g, "<br/>")
    .replace(/\n/g, "<br/>")
    .replace(/\s/g, " ");
}

// 检测用户名是否合格
function checkloginname(str) {
  var rule = new RegExp("^[a-zA-z0-9_]{5,}$");
  if (rule.test(str)) {
    return true;
  } else return false;
}
// 检测密码是否合格
function checkpassword(str) {
  var rule = new RegExp("^S{6,}$");
  if (rule.test(str)) {
    return true;
  } else return false;
}
// 检测邮箱是否合格
// 校验电子邮件
function checkEmailone(cstr) {
  var Rule = new RegExp(
    "^[A-Za-z0-9_-]+@[A-Za-z0-9-]+.(com|cn|com.cn|net|net.cn)$"
  );
  if (Rule.test(cstr)) {
    return true;
  } else {
    return false;
  }
}

// 检则是否是汉字组合
function checkChinese(cstr) {
  var rule = new RegExp("^[\u4e00-\u9fa5]+$");
  if (rule.test(cstr)) {
    return true;
  } else {
    return false;
  }
}

// 检则QQ号码是否正确
function checkQQnum(cstr) {
  var rule = new RegExp("^[1-9][0-9]{4,}$");
  if (rule.test(cstr)) {
    return true; // 是QQ号
  } else return false; // 不是QQ号码
}

// 检则是否是数字
function checkIsnum(num) {
  var rule = new RegExp("^[0-9]{1,}$");
  if (rule.test(num)) {
    return true; // 是数字
  } else return false; // 不是数字
}
// 判断是不是钱
function checkIsMoney(num) {
  var rule = new RegExp("^[0-9]{1}.[0-9]{1,2}$|^[0-9]{1,}$");
  if (rule.test(num)) {
    return true; // 是金钱
  } else return false; // 不是数字
}
// 检则是否为日期
function checkIsDate(MyDate) {
  var rule = new RegExp(
    "^2[0-9]{3}-([1-9]|0[0-9]|1[0-2])-([1-9]|1[0-9]|2[0-9]|3[0-1])$"
  );
  if (rule.test(MyDate)) {
    return true; // 是数字
  } else return false; // 不是数字
}

//rsa加密函数 后端php在解密
function encryptionRsa(value) {
  //js前端rsa加密 密码password加密
  var rsa = new RSAKey();
  //openssl生成的modulus,十六进制数据
  var modulus =
    "CBB672EFA1D9C6E4FCA8942C907C7D55C8B469BB1ADDFD34912154737987CD4631F75A32F38AEEEC06CD0F067A321646F33CD0B78189A09BA39DA5A8947BDF7F7DDE0F8FD8993F9EB7F2283D15367534933321EF5B7D66F217C26D162C692578D186F494E26AAE73F7A088EFC89D3AD7D2D8692BCDC2A7D3F1213BA5D847C03D";
  //openssl生成秘钥时的e的值(0x10001)
  var exponent = "10001";
  rsa.setPublic(modulus, exponent);
  return rsa.encrypt(value); //前端rsa加密后的
}

/**
 * 去除空格 两边 和 字符串中间空格
 * @param string  string  字符串内容
 * @return 去掉空格后的字符串
 */
function removeSpace(str) {
  return (str = str.replace(/\s+/g, ""));
}

/**
 * 获取元素距离左边屏幕的像素
 * @param  {[type]} obj [description]
 * @return {[type]}     [description]
 */
function getAbsLeft(obj) {
  var l = obj.offsetLeft;
  while (obj.offsetParent != null) {
    obj = obj.offsetParent;
    l += obj.offsetLeft;
  }
  return l;
}

//js获取url传递参数，js获取url？号后面的参数
function GetRequest() {
  var url = location.search; //获取url中"?"符后的字串
  if (url != "" && url != null) {
    var theRequest = new Object();
    if (url.indexOf("?") != -1) {
      var str = url.substr(1);
      strs = str.split("&");
      for (var i = 0; i < strs.length; i++) {
        theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
      }
    }
    return theRequest;
  } else {
    return 0;
  }
}

//时间戳转换成日期格式 小时 分钟 05:06
function timeConver(timeStamp) {
  var newTime = new Date(timeStamp);
  var Hours =
    newTime.getHours() > 9 ? newTime.getHours() : "0" + newTime.getHours();
  var Minutes =
    newTime.getMinutes() > 9
      ? newTime.getMinutes()
      : "0" + newTime.getMinutes();
  return Hours + ":" + Minutes;
}

//获取时间戳
function timestampToTime(timestamp) {
  var date = new Date(timestamp); //时间戳为10位需*1000，时间戳为13位的话不需乘1000
  // var date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
  Y = date.getFullYear() + "-";
  M =
    (date.getMonth() + 1 < 10
      ? "0" + (date.getMonth() + 1)
      : date.getMonth() + 1) + "-";
  D = date.getDate() + " ";
  h = date.getHours() + ":";
  m = date.getMinutes() + ":";
  s = date.getSeconds();
  return Y + M + D + h + m + s;
}
