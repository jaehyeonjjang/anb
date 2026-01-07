var _saveFlag = false;
var _syncId = 0;

function saveProcess(func) {
    if ($('#detailModal').is(':visible')) {
        if (_filetype == FILETYPE_FLOORITEM) {
            var cnt = 0;
            for (var i = 0; i < 1000; i++) {
                var name = $('#name' + i).val();
                var imagename = $('#imagename' + i).val();

                if (name != '' & imagename != '')
                    cnt++;
                else
                    break;
            }

            if (_items.length < cnt) {
                for (var i = 0; i < cnt - _items.length; i++) {
                    var item = {
                        mode: 20,
                        type: 20,
                        points: null,
                        x: x,
                        y: y,
                        color:'',
                        selected: false,
                        group: 0,
                        number: 0,

                        name: $('#name' + i).val(),
                        fault: $('#fault' + i).val(),
                        content: $('#content' + i).val(),
                        report: $('#report' + i).val(),
                        usermemo: $('#usermemo' + i).val(),
                        aptmemo: $('#aptmemo' + i).val(),
                        width: $('#width' + i).val(),
                        length: $('#length' + i).val(),
                        count: $('#count' + i).val(),
                        progress: $('#progress' + i).val(),
                        remark: $('#remark' + i).val(),
                        imagename: $('#imagename' + i).val(),
                        filename: ''
                    };

                    _items.push(item);
                }
            } else if (_items.length > cnt) {
                // 삭제해야함
            }
        }
        
        if (_filetype == 1 || _filetype == 5 || _filetype == 0 || _filetype == FILETYPE_FLOORITEM) {
            for (var i = 0; i < _items.length; i++) {
                if (!isNumber(_items[i].mode))
                    continue;

                _items[i].name = $('#name' + i).val();
                _items[i].fault = $('#fault' + i).val();
                _items[i].content = $('#content' + i).val();
                _items[i].report = $('#report' + i).val();
                _items[i].usermemo = $('#usermemo' + i).val();
                _items[i].aptmemo = $('#aptmemo' + i).val();
                _items[i].width = $('#width' + i).val();
                _items[i].length = $('#length' + i).val();
                _items[i].count = $('#count' + i).val();
                _items[i].progress = $('#progress' + i).val();
                _items[i].remark = $('#remark' + i).val();
                _items[i].imagename = $('#imagename' + i).val();
            }
        }
    }
    

    for (var i = 0; i < _items.length; i++) {
        _items[i].selected = false;
    }
    
    if (_filetype == 5) {
        for (var i = 0; i < _items.length; i++) {
            _items[i].name = $('#cboFloorName').val();
            _items[i].memo = {type: _floortypeId};
        }
    }

    if (_filetype == 3) {
        makeItems2();

        updateCalc();
    }

    var send = [];
    for (var i = 0; i < _items.length; i++) {
        var item = {
            apt: _apt,
            image: parseInt(_image),
            imagetype: _filetype,
            type: parseInt(_items[i].mode),
            mode: parseInt(_items[i].mode),
            x: _items[i].x,
            y: _items[i].y,
            point: JSON.stringify(_items[i].points),
            number: parseInt(_items[i].number),
            group: _items[i].group,
            name: _items[i].name,
            fault: _items[i].fault,
            content: _items[i].content,
            report: parseInt(_items[i].report),
            usermemo: _items[i].usermemo,
            aptmemo: _items[i].aptmemo,
            width: parseFloat(_items[i].width),
            length: parseFloat(_items[i].length),
            count: _items[i].count,
            progress: _items[i].progress,
            remark: _items[i].remark,
            imagename: _items[i].imagename,
            filename: _items[i].filename,
            memo: JSON.stringify(_items[i].memo)
        };

        if (_filetype == 0) {
            if (item.name == '') {
                continue;
            }
        }

        send.push(item);
    }

    var item = {
        apt: _apt,
        image: parseInt(_image),
        imagetype: _filetype,
        type: 100,
        x: 0,
        y: 0,
        point: '',
        number: 0,
        group: 0,
        name: '',
        fault: '',
        content: '' + _imageZoom,
        report: 1,
        usermemo: '',
        aptmemo: '',
        width: 0.0,
        length: 0.0,
        count: '',
        progress: '',
        remark: '',
        imagename: '',
        filename: '',
        memo: '' + _arrowZoom,
    };

    send.push(item);

    if (_filetype == 3) {
        for (var i = 0; i < 10; i++) {
            var item = {
                apt: _apt,
                image: parseInt(_image),
                imagetype: _filetype,
                type: 200,
                name: $('#flow' + i).val(),
                content: $('#sh1' + i).val(),
                width: parseFloat($('#sh2' + i).val()),
                length: parseFloat($('#n' + i).val()),
                count: $('#ps' + i).val()
            };

            send.push(item);
        }
    }

    if (_serverMode == 'mobile' || _serverMode == 'electron') {
        makeGroupNumber();
        var max = 0;
        var items = [];
        var datas = JSON.parse(localStorage.getItem('datas'));
        if (datas.datas != null) {
            for (var i = 0; i < datas.datas.length; i++) {
                var item = datas.datas[i];

                if (item.Id > max)
                    max = item.Id;

                if (parseInt(item.Image) == parseInt(_image))
                    continue;

                items.push(item);
            }
        }

        var d = new Date();
        var date = d.getFullYear() + '-' + pad(d.getMonth() + 1, 2) + '-' + pad(d.getDate(), 2) + ' ' + pad(d.getHours(), 2) + ':' + pad(d.getMinutes(), 2) + ':' + pad(d.getSeconds(), 2);

        
        for (var i = 0; i < send.length; i++) {
            max++;

            var s = send[i];

            var item = {
                Id: max,
                Apt: _apt,
                Content: s.content,
                Report: s.report,
                Usermemo: s.usermemo,
                Aptmemo: s.aptmemo,
                Count: s.count,
                Date: date,
                Extra: '',
                Filename: s.filename,
                Group: s.group,
                Image: s.image,
                Imagename: s.imagename,
                Imagetype: s.imagetype,
                Length: s.length,
                Memo: s.memo,
                Name: s.name,
                Fault: s.fault,
                Number: s.number,
                Point: s.point,
                Progress: s.progress,
                Remark: s.remark,
                Type: s.type,
                User: _user.Id,
                Width: s.width,
                X: s.x,
                Y: s.y
            }

            items.push(item);
        }

        datas.datas = items;

        _syncId++;
        datas.syncs.push({
            Id: _syncId,
            Image: parseInt(_image),
            Date: date 
        });

        localStorage.setItem('datas', JSON.stringify(datas));

        _saveFlag = false;
        $('#divWaitPopup').hide();
        initModified();

        /*
          if (func == null)
          alert('저장되었습니다');
          else
          func();
        */


        var img = canvas.toDataURL();

        var filename = _apt + '-' + _image + '.png';
        saveFile(filename, img, function(err, filename) {
            _saveFlag = false;
            $('#divWaitPopup').hide();
            initModified();

            if (err != null)
                alert(err);
            if (func == null)
                alert('저장되었습니다');
            else
                func();
        });
        /*
          var filename = getImageName('webdata/' + _apt + '-' + _image + '.png');
          saveFile(filename, img, function() {
          _saveFlag = false;
          $('#divWaitPopup').hide();
          initModified();

          if (func == null)
          alert('저장되었습니다');
          else
          func();
          });
        */
    } else {
        $.ajax({
            url: _serverAddr + '/api/data/multiinsert',
            type: 'POST',
            data: JSON.stringify({image: parseInt(_image), data: send}),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                if (_filetype == FILETYPE_FLOORITEM) {
                    _saveFlag = false;
                    $('#divWaitPopup').hide();
                    initModified();

                    redraw();

                    if (func == null) {
                        alert('저장되었습니다');
                    } else {
                        func();
                    }

                    return;
                }
                
                
                var img = canvas.toDataURL();
                params = {
                    apt: _apt,
                    image: _image,
                    data: canvas.toDataURL()
                };
                $.post(_serverAddr + '/api/data/image', params, function(data) {
                    _saveFlag = false;
                    $('#divWaitPopup').hide();
                    initModified();

                    redraw();

                    if (func == null) {
                        alert('저장되었습니다');
                    } else {
                        func();
                    }
                });
            }
        }).fail(function() {
            _saveFlag = false;
            $('#divWaitPopup').hide();
            initModified();
            alert('서버에 접속할수 없습니다. 잠시후 다시 이용해주세요');
        });
    }
}

function save(func) {
    if (_saveFlag == true) {
        alert('이미 저장중 중입니다');
        return;
    }

    closeSinglePopup();

    _saveFlag = true;
    $('#divWaitPopup').show();

    initUndo();
    defaultZoom();

    setTimeout(function() {
        saveProcess(func);
    }, 500);
}
