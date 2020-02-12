$("#input").submit(function(event) {
    event.preventDefault();

    $longURL = $(this).find("input[name='long-url']").val()

    $.ajax({
        url: "/add",
        type: "get",
        data: {
            longURL: $longURL
        },
        contentType: "application/json; charset=utf-8",
        success: function(response) {
            $("#result").first().fadeIn("slow");
            $("input[name='short-url']").first().val(response);
        },
        error: function(xhr) {
            alert("An error occured...");
        }
    });
});