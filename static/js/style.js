document.addEventListener("DOMContentLoaded", function() {
    console.log("DOM fully loaded and parsed");
    var gifts = document.querySelectorAll('.gift img');
    console.log("Number of images found: ", gifts.length);
    gifts.forEach(function(img) {
        var randomClass = Math.random() > 0.5 ? 'image-offset-right-bottom' : 'image-offset-left-top';
        console.log("Applying class: ", randomClass);
        img.classList.add(randomClass);
    });
});

window.addEventListener("wheel", function(e) {
    console.log("Wheel event detected");
    if (e.deltaX === 0) {
        e.stopPropagation();
        e.preventDefault();
        console.log("Scrolling horizontally by: ", e.deltaY);
        window.scrollBy(e.deltaY, 0);
    }
});
