$(document).ready(function() {
    const urlUpload = 'https://dev.konkov.xyz/api/upload';
    const urlDownload = 'https://dev.konkov.xyz/api/download';

    $('#uploadButton').click(function(e) {
        e.preventDefault();
    
        var formData = new FormData();
        formData.append('source', $('#inputForm').prop('files')[0]);

        success = function(data, textStatus, jqXHR) {
            console.log(data, textStatus, jqXHR);
            $('#token').html(data['token_id'])
        }
        
        error = function(result) {
            console.log(result)
            alert('error');
        }
    
        $.ajax({
            type: 'POST',
            url: urlUpload,
            data: formData,
            processData: false,
            contentType: false,
            success: success,
            error: error
        });
    })

    $('#downloadButton').click(function(e) {
        e.preventDefault();
        
        urlWithParameters = urlDownload + '?token_id=' + $('#floatingInput').val()

        success = function(data, textStatus, jqXHR) {
            console.log(data, textStatus, jqXHR)

            var blob=new Blob([atob(data.data)]);
            var link=document.createElement('a');
            link.href=window.URL.createObjectURL(blob);
            link.download=data.file_name;
            link.click();
        }

        error = function(result) {
            console.log(result)
            alert('error');
        }

        $.ajax({
            type: 'GET',
            url: urlWithParameters,
            data: null,
            processData: false,
            contentType: false,
            success: success,
            error: error
        });
    })
 });