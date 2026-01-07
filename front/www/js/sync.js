var _syncFlag = false;

function sync() {
    if (_syncFlag == true) {
        alert('이미 동기화 중입니다');
        return;
    }

    if (!confirm('동기화 하시겠습니까\n동기화 이후 자동 로그아웃되며 온라인 상태에서 새로 로그인해야 재작업이 가능합니다'))
        return;

    if (_saveFlag == true) {
        alert('데이터 저장중입니다. 잠시후 다시 이용해 주세요');
        return;
    }

    if (isModified() == true) {
        if (!confirm('기존에 작업한 내역이 저장되지 않았습니다.\n저장없이 동기화 시키면 해당 내용은 반영되지 않습니다.\n동기화하시겠습니까'))
            return;
    }

    closeSinglePopup();
    
    _syncFlag = true;
    $('#divWaitPopup').show();

    if (_serverMode == 'mobile') {
        var lists = JSON.parse(localStorage.getItem('lists'));            
        var datas = JSON.parse(localStorage.getItem('datas'));

        var images = [];
        if (lists.images != null) {
            for (var i = 0; i < lists.images.length; i++) {
                var item = lists.images[i];

                if (item.Id < 0) {
                    images.push(item);
                }
            }
        }

        var imagefloors = [];
        if (lists.imagefloors != null) {
            for (var i = 0; i < lists.imagefloors.length; i++) {
                var item = lists.imagefloors[i];

                if (item.Id < 0) {
                    imagefloors.push(item);
                }
            }
        }

        var fns = [];
        
        var items = {
            Datas: datas.datas,
            Syncs: datas.syncs,
            Images: images,
            Imagefloors: imagefloors
        }


        console.log('+++++++++++++++++++++++++++++++++++++++++++');
        console.log(datas.syncs);
        console.log('+++++++++++++++++++++++++++++++++++++++++++');
        var imgs = {};
        if (datas.syncs != null) {
            for (var i = 0; i < datas.syncs.length; i++) {
                var item = datas.syncs[i];

                imgs[item.Image] = item.Image;
            }
        }

        for (key in imgs) {
            fns.push(uploadCanvas(_apt, key));

            if (datas.datas != null) {
                
                for (var i = 0; i < datas.datas.length; i++) {
                    var item = datas.datas[i];

                    if (item.Image != key) {
                        continue;
                    }
                    
                    if (item.Filename != undefined && item.Filename != null && item.Filename != '') {

                        console.log('item.Filename :' + item.Filename);
                        fns.push(uploadCameraImage(item.Filename));

                    }
                }
            }
        }
        
        
        $.ajax({
            url: _serverAddr + "/api/sync/upload",
            type: "post",
            accept: "application/json",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(items),
            dataType: "json",
            success: function(data) {
                $.when.apply($, fns).done(function() {
                    $.post(_serverAddr + '/api/sync/complete', 'image=' +  _image, function(data) {
                    });
                           
                    _syncFlag = false;
                    $('#divWaitPopup').hide();                            
                    
                    alert('동기화가 완료되었습니다.');

                    clickLogout(false);
                })
                 .fail(function() {
                     _syncFlag = false;
                     $('#divWaitPopup').hide();

                     alert('동기화에 실패했습니다');                 
                 });            
            },                    
            error: function(jqXHR,textStatus,errorThrown) {
                _syncFlag = false;
                $('#divWaitPopup').hide();

                alert('동기화에 실패했습니다');                     
            }
        });
    } else {
        $.post(_serverAddr + '/api/sync/sync', 'image=' +  _image, function(data) {
            $.post(_serverAddr + '/api/sync/complete', 'image=' +  _image, function(data) {
            });
            
            _syncFlag = false;
            $('#divWaitPopup').hide();
            alert('동기화가 완료되었습니다.');
        }).fail(function() {
            _syncFlag = false;
            $('#divWaitPopup').hide();
            alert('서버에 접속할수 없습니다. 잠시후 다시 이용해주세요');
        });            
    }               

}

