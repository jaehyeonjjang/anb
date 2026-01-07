var _downloadFlag = false;

function clickDownload() {
    if (_downloadFlag == true) {
        alert('이미 다운로드 중입니다');
        return;
    }

    if (!confirm('데이터를 다운로드 하시겠습니까'))
        return;

    _downloadFlag = true;
    $('#divWaitPopup').show();

    if (_serverMode == 'mobile' || _serverMode == 'electron') {
        $.post(_serverAddr + '/api/json/list', '', function(lists) {
            console.log(lists);

            $.post(_serverAddr + '/api/json/data', '', function(datas) {
                console.log(datas);

                var fns = [];

                localStorage.setItem('lists', JSON.stringify(lists));

                for (var i = 0; i < lists.images.length; i++) {
                    var item = lists.images[i];

                    if (item.Filename == '')
                        continue;


                    if (_serverMode == 'mobile') {
                        console.log('download image : ' + _serverAddr + '/' + item.Filename);
                        fns.push(downloadFileToDevice(_serverAddr + '/' + item.Filename, getImageName(item.Filename)));
                    }
                }

                for (var i = 0; i < lists.statuss.length; i++) {
                    var item = lists.statuss[i];

                    if (item.Type != 8)
                        continue;

                    if (item.Content == '')
                        continue;

                    if (_serverMode == 'mobile') {
                        console.log('download status : ' + _serverAddr + '/' + item.Content);
                        fns.push(downloadFileToDevice(_serverAddr + '/' + item.Content, getImageName(item.Content)));
                    }
                }

                datas.syncs = [];

                localStorage.setItem('datas', JSON.stringify(datas));

                if (datas.datas != null) {
                    for (var i = 0; i < datas.datas.length; i++) {
                        var item = datas.datas[i];

                        if (item.Filename == '')
                            continue;

                        if (_serverMode == 'mobile') {
                            console.log('data image status : ' + _serverAddr + '/' + item.Filename);
                            fns.push(downloadFileToDevice(_serverAddr + '/' + item.Filename, getImageName(item.Filename)));
                        }
                    }
                }

                $.when.apply($, fns).done(function() {
                    _downloadFlag = false;
                    $('#divWaitPopup').hide();

                    alert('다운로드가 완료되었습니다.');
                })
                 .fail(function() {
                     _downloadFlag = false;
                     $('#divWaitPopup').hide();

                     alert('다운로드에 실패했습니다');
                 });

            });
        });            
    } else { 
        $.post(_serverAddr + '/api/sync/download', '', function(data) {
            _downloadFlag = false;
            $('#divWaitPopup').hide();

            alert('다운로드가 완료되었습니다.');
        });
    }
}

function saveFile(filename, DataBlob, callback) {
    //console.log(cordova.file);
    var target = cordova.file.dataDirectory;

    window.resolveLocalFileSystemURL(target, function (dir) {
        dir.getFile(filename, { create: true }, function (file) {
            file.createWriter(function (fileWriter) {
                fileWriter.write(DataBlob)
                callback(null, filename);
            }, function (err) {
                callback(err);
            })
        })
    }, function(err) {
        console.log(err);
    })
}

function downloadFileToDevice (fileurl, filename) {
    var deferred = $.Deferred();

    var blob = null
    var xhr = new XMLHttpRequest()
    xhr.open('GET', fileurl)
    xhr.responseType = 'blob' // force the HTTP response, response-type header to be blob
    xhr.onload = function () {
        //console.log('onLoad : ' + filename);
        blob = xhr.response; // xhr.response is now a blob object
        saveFile(filename, xhr.response, function(err, filePath)  {
            if (err) {
                console.log('An error was found: ', err);
                deferred.reject(err);
            } else {
                //console.log('file downloaded successfully to: ' + filePath);
                deferred.resolve(filePath);
            }
        });
    };

    xhr.onerror = function(err) {
        console.log('An error was found: ', err);
        deferred.reject(err);
    };

    try {
        xhr.send();
    } catch (e) {
        deferred.reject(err);
    }

    //return deferred.promise();
    return deferred;
}



function clickAptDownload(id, title) {
    if (_downloadFlag == true) {
        alert('이미 다운로드 중입니다');
        return;
    }

    if (!confirm('데이터를 다운로드 하시겠습니까'))
        return;

    closeShowPopup();

    _downloadFlag = true;
    $('#divWaitPopup').show();

    
    $.post(_serverAddr + '/api/json/list', 'apt=' + id, function(lists) {
        console.log(lists);

        $.post(_serverAddr + '/api/json/data', 'apt=' + id, function(datas) {
            console.log(datas);

            var fns = [];

            localStorage.setItem('lists', JSON.stringify(lists));

            for (var i = 0; i < lists.images.length; i++) {
                var item = lists.images[i];

                if (item.Apt != id)
                    continue;

                if (item.Filename == '')
                    continue;


                if (_serverMode == 'mobile') {
                    console.log('download image : ' + _serverAddr + '/' + item.Filename);
                    fns.push(downloadFileToDevice(_serverAddr + '/' + item.Filename, getImageName(item.Filename)));
                }
            }

            for (var i = 0; i < lists.statuss.length; i++) {
                var item = lists.statuss[i];

                if (item.Type != 8)
                    continue;

                if (item.Content == '')
                    continue;

                if (_serverMode == 'mobile') {
                    console.log('download status : ' + _serverAddr + '/' + item.Content);
                    fns.push(downloadFileToDevice(_serverAddr + '/' + item.Content, getImageName(item.Content)));
                }
            }

            datas.syncs = [];

            localStorage.setItem('datas', JSON.stringify(datas));

            if (datas.datas != null) {
                for (var i = 0; i < datas.datas.length; i++) {
                    var item = datas.datas[i];

                    if (item.Apt != id)
                        continue;
                    
                    if (item.Filename == '')
                        continue;

                    if (_serverMode == 'mobile') {
                        console.log('data image status : ' + _serverAddr + '/' + item.Filename);
                        fns.push(downloadFileToDevice(_serverAddr + '/' + item.Filename, getImageName(item.Filename)));
                    }
                }
            }

            $.when.apply($, fns).done(function() {
                _downloadFlag = false;
                $('#divWaitPopup').hide();

                //alert('다운로드가 완료되었습니다.');

                clickApt(id, title);
            })
             .fail(function() {
                 _downloadFlag = false;
                 $('#divWaitPopup').hide();

                 alert('다운로드에 실패했습니다');
             });

        });
    });
}
