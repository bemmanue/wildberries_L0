const form = $("#order-form");

$(document).ready(function() {

    jQuery.validator.addMethod("alphanumeric", function(value, element) {
        return this.optional(element) || /^[a-zA-Z\d]+$/.test(value);
    }, "Only alphanumeric characters");

    $("form").validate({
        rules: {
            username: {
                required: true,
                alphanumeric: true
            }
        }
    })

    $(document).on('submit', form, function(e){
        let output = document.getElementById("order-data")

        e.preventDefault();
        $.ajax({
            type: "GET",
            url: "/order/" + $('#order_uid').val(),
            success: function (data, textStatus, jqXHR) {
                output.innerText = JSON.stringify(JSON.parse(jqXHR.responseText), null, 2)
            },
            error: function () {
                output.innerText = "not found"
            },
        });
    });
});