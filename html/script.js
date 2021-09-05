$(document).ready(function() {
    const urlUpload = 'https://dev.konkov.xyz/api/file';
    const urlDownload = 'https://dev.konkov.xyz/api/file';

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

            // check for a filename
            var filename = "";
            var disposition = jqXHR.getResponseHeader('Content-Disposition');
            if (disposition && disposition.indexOf('attachment') !== -1) {
                var filenameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/;
                var matches = filenameRegex.exec(disposition);
                if (matches != null && matches[1]) filename = matches[1].replace(/['"]/g, '');
            }
            
            var link=document.createElement('a');
            link.href=window.URL.createObjectURL(data);
            link.download=filename;
            link.click();
        }

        error = function(result) {
            console.log(result)
            alert('error');
        }

        $.ajax({
            type: 'GET',
            url: urlWithParameters,
            xhrFields: {
                responseType: 'blob' 
            },
            data: null,
            processData: false,
            contentType: false,
            success: success,
            error: error
        });
    })
 });