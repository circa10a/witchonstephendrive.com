canvas = document.getElementById("canvas");
var context = canvas.getContext('2d');
W = canvas.width = window.innerWidth;
H = canvas.height = window.innerHeight;
generatorStock = [];

generatorStock.push(new particleGenerator(2, 3, 30, 30));

function randomInt(min, max) {
    return Math.floor(min + Math.random() * (max - min + 1));
}

function particle(vx, vy, size) {
    this.radius = randomInt(0, size);
    this.x = W / 2;
    this.y = H / 2;
    this.alpha = 1;
    this.vx = randomInt(-vx, vx);
    if (Math.random() < 0.1) {
        this.vy = randomInt(-vy, -vy - 3);
    }
    else {
        this.vy = randomInt(0, -vy);
    }
    this.lifetime = 100;
}
particle.prototype.update = function () {
    this.lifetime -= 2;
    this.x += this.vx;
    this.y += this.vy;

    if (this.lifetime < 20) {
        this.radius -= 2;
    }

    context.beginPath();
    context.arc(this.x, this.y, this.radius, Math.PI * 2, false);
    context.fillStyle = "rgba(255,255,255,0.5)";
    context.strokeStyle = "#2f2";
    context.lineWidth = 10;
    context.fill();
    context.stroke();
    context.closePath();

}

function particleGenerator(vx, vy, size, maxParticles) {
    this.size = size;
    this.vx = vx;
    this.vy = vy;
    this.maxParticles = maxParticles;
    this.particles = [];
}
var freq = 0.5;
particleGenerator.prototype.animate = function () {
    if (this.particles.length < this.maxParticles && Math.random() < freq) {
        this.particles.push(new particle(this.vx, this.vy, this.size));
        if (this.particles.length == this.maxParticles / 2) {
            freq = 0.1;
        }
    }
    for (var i = 0; i < this.particles.length; i++) {
        p = this.particles[i];
        p.update();
        if (p.radius < 10) {
            this.particles[i] = new particle(this.vx, this.vy, this.size);
        }
    }
}

function update() {
    context.clearRect(0, 0, W, H);

    for (var i = 0; i < generatorStock.length; i++) {
        generatorStock[i].animate();
    }
    window.requestAnimationFrame(update);
}

update();

window.addEventListener('resize', function (e) {
    W = window.innerWidth;
    H = window.innerHeight;
}, false);
