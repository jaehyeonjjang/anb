function press_it() {
    if(event.keyCode == 13) {
        login();
    }
}

function checkLogin(data) {
    if (data.code == 'user not found') {
        alert('로그인 정보가 정확하지 않습니다');
    } else if (data.code == 'wrong password') {
        alert('로그인 정보가 정확하지 않습니다');
    } else if (data.code == 'not permit') {
        alert('관리자 승인 대기 중입니다. 관리자 승인이후 사용 가능합니다.');
    } else {
        if (_serverMode == 'mobile' || _serverMode == 'electron') {
            localStorage.setItem('session', JSON.stringify(data.user));        
        }

        if ($('#chkSave').is(':checked') == true) {
            localStorage.setItem('saveid', 'y');
            localStorage.setItem('saveloginid', $('#txtLoginid').val());
            localStorage.setItem('savepasswd', $('#txtPasswd').val());
        } else {
            localStorage.setItem('saveid', 'n');
            localStorage.setItem('saveloginid', '');
            localStorage.setItem('savepasswd', '');
        }
        
        _user = data.user;
        showAptPopup();
    }    
}

var _user = null;


function login() {
    var f = document.login_form;
    if (f.loginid.value == '') {
        alert('로그인 아이디를 입력하세요');
        f.loginid.focus();
        return false;
    }

    if (f.passwd.value == '') {
        alert('비밀번호를 입력하세요');
        f.passwd.focus();
        return false;
    }

    var params = $('#login_form').serialize();

    /*
      if (_serverMode == 'mobile' || _serverMode == 'electron') {
      var loginid = f.loginid.value;
      var passwd = SHA256(f.passwd.value);

      var code = 'user not found';

      var lists = JSON.parse(localStorage.getItem('lists'));
      if (lists == null) {
      alert('데이터를 먼저 다운로드 받으세요');
      return;
      }

      for (var i = 0; i < lists.users.length; i++) {
      var item = lists.users[i];

      if (item.Loginid == loginid) {
      if (item.Passwd == passwd) {
      code = 'ok';

      _user = item;
      
      localStorage.setItem('session', JSON.stringify(item));
      } else {
      code = 'wrong password';                    
      }
      break;
      }
      }
      
      checkLogin({code: code, user: _user});
      } else {
      $.post(_serverAddr + '/api/login/login', params, function(data) {
      checkLogin(data);
      }).fail(function() {
      alert('서버에 접속할수 없습니다. 잠시후 다시 이용해주세요');
      });
      }
    */

    $.post(_serverAddr + '/api/login/login', params, function(data) {
        checkLogin(data);
    }).fail(function() {
        alert('서버에 접속할수 없습니다. 잠시후 다시 이용해주세요');
    });        

    return false;
}
