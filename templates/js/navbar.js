// Select nav elements
const burger = document.querySelector('.burger i');
const nav = document.querySelector('.nav');

//trigers style change when burger is pressed
function toggleNav() {
    burger.classList.toggle('fa-bars');
    burger.classList.toggle('fa-times');
    nav.classList.toggle('nav-active');
}

// burger event listeners
burger.addEventListener('click', function() {
    toggleNav();
});

