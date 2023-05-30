//封装element message
function message(message, type, callback) {
  // Vue.prototype.$message({
  ElMessage({
    showClose: true,
    message: message,
    type: type,
    onClose: () => {
      if (callback) {
        callback();
      }
    },
  });
}

function baseConfirm(content, title, callback1, callback2) {
  Vue.prototype
    .$confirm(content, title || "温馨提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      closeOnClickModal: false,
      type: "warning",
    })
    .then(() => {
      if (callback1) {
        callback1();
      }
    })
    .catch(() => {
      if (callback2) {
        callback2();
      }
    });
}

// ajax 返回操作封装
function response(res, callback) {
  const { code, msg, data } = res;
  if (code == 1) {
    // if (callback && data) {
    if (callback) {
      callback(data);
    }
  } else {
    message(msg, "error");
    // layer.msg(msg);
    // baseMessage(msg, 'error');
  }
}

//验证用户名  element 验证器=====================================================================================================
function checkUsername(rule, value, callback) {
  var pattern = new RegExp("^[a-zA-z0-9_]{5,50}$");
  if (pattern.test(value) == false) {
    callback(new Error("用户名只支持数字字母或下划线"));
  } else {
    callback();
  }
}

//验证验证码长度
function checkCodeLength(rule, value, callback) {
  if (value.length != 4) {
    callback(new Error("请输入4位数的验证码"));
  } else {
    callback();
  }
}

//验证手机号格式
function checkUserphone(rule, value, callback) {
  var pattern = /^1[345678]\d{9}$/;
  if (pattern.test(value) == false) {
    callback(new Error("请输入正确的手机号码"));
  } else {
    callback();
  }
}

//验证邮箱
function checkEmailFormat(rule, value, callback) {
  var pattern = /^(\w)+(\.\w+)*@(\w)+((\.\w{2,3}){1,3})$/;
  if (pattern.test(value) == false) {
    callback(new Error("请输入正确的邮箱"));
  } else {
    callback();
  }
}

//验证正整数
function isNum(rule, value, callback) {
  if (!checkIsnum(value)) {
    callback(new Error("只能是正整数"));
  } else {
    callback();
  }
}

// ip判断
function isIP(rule, value, callback) {
  if (!checkIP(value)) {
    callback(new Error("请输入合法的ip"));
  } else {
    callback();
  }
}

// ===================================================================================================================================

// 检测用户名是否合格
function checkloginname(str) {
  var rule = new RegExp("^[a-zA-z0-9_]{5,50}$");
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

// 检测ip
function checkIP(str) {
  const reg = /^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$/;
  return reg.test(str);
}
