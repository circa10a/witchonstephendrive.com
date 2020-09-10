'use strict';
var Ghost = function () {
    function Ghost(svg, start, end) {
        this.id = String(Math.round(Math.random() * 999999999999999));
        this.group = svg.group();
        this.startPoint = start;
        this.endPoint = end;
        this.startThickness = 0;
        this.endThickness = 150 + Math.round(Math.random() * 50);
        this.guidePosition = Math.random() * 1000;
        this.frequency = 0.01 + Math.random() * 0.01;
        this.amplitude = 20 + Math.random() * 40;
        this.height = 0;
        this.endHeight = 150 + Math.round(Math.random() * 100);
        this.y = 0;
        var faceAttr = {
            fill: '#111111',
            opacity: 0.9,
            stroke: 'none'
        };

        this.body = this.group.path().attr({
            fill: '#eeeeee',
            opacity: 0.8,
            stroke: 'none'
        });

        this.eyeLeft = this.group.path().attr(faceAttr);
        this.eyeRight = this.group.path().attr(faceAttr);
        this.mouth = this.group.path().attr(faceAttr);
        this.updateGuide();
    }
    Ghost.prototype.remove = function () {
        this.group.remove();
    };
    Ghost.prototype.updateGuide = function () {
        this.guide = [];
        var height = this.startPoint.y - this.endPoint.y;
        var widthChange = this.startPoint.x - this.endPoint.x;
        var y = this.startPoint.y;
        while (y-- >= this.endPoint.y) {
            var x = this.startPoint.x + (widthChange - widthChange / height * y);
            var wave = Math.sin(y * this.frequency + this.guidePosition);
            this.guide.push({ y: y, x: x + (wave * this.amplitude / 2 + this.amplitude / 2) });
        }
        //console.log(this.guide)
    };
    Ghost.prototype.start = function (onComplete) {
        TweenMax.to(this, 2, { y: this.guide.length, height: this.endHeight, position: '+=6', ease: SlowMo.ease.config(0.3, 0.3, false), onComplete: onComplete, onCompleteParams: [this] });
    };
    Ghost.prototype.getPointAlongGuide = function (y, offsetXPercentage) {
        if (this.guide.length) {
            if (y >= this.guide.length)
                y = this.guide.length - 1;
            if (y < 0)
                y = 0;
            var thicknessDifference = this.endThickness - this.startThickness;
            var percentageAlongGuide = y / this.guide.length * 100;
            var thickness = this.startThickness + thicknessDifference / 100 * percentageAlongGuide;
            var xOffset = thickness / 2 / 100 * offsetXPercentage;
            return { x: this.guide[y].x + xOffset, y: this.guide[y].y };
        }
        return { x: 0, y: 0 };
    };
    Ghost.prototype.drawPath = function (pathPoints) {
        var points = [];
        for (var i = 0; i < pathPoints.length; i++) {
            var subPoints = [];
            for (var j = 0; j < pathPoints[i].points.length / 2; j++) {
                var p = pathPoints[i].points.slice(j * 2, j * 2 + 2);
                //console.log(i, p)
                var point = this.getPointAlongGuide(Math.round(p[1]), p[0]);
                subPoints.push(point.x);
                subPoints.push(point.y);
            }
            points.push(pathPoints[i].type + subPoints.join(' '));
        }
        return points.join(' ') + 'Z';
    };
    Ghost.prototype.draw = function () {
        if (this.height > 0) {
            var y = Math.round(this.y);
            var height = Math.round(this.height);
            var heightChunks = height / 6;
            var body = [
                { type: 'M', points: [10, y] },
                { type: 'Q', points: [75, y, 80, y - heightChunks * 2] },
                {
                    type: 'L', points: [85, y - heightChunks * 3,
                        90, y - heightChunks * 4,
                        95, y - heightChunks * 5,
                        100, y - heightChunks * 6,
                        75, y - heightChunks * 5,
                        50, y - heightChunks * 6,
                        25, y - heightChunks * 5,
                        0, y - heightChunks * 6,
                        -25, y - heightChunks * 5,
                        -50, y - heightChunks * 6,
                        -75, y - heightChunks * 5,
                        -100, y - heightChunks * 6,
                        -95, y - heightChunks * 5,
                        -90, y - heightChunks * 4,
                        -85, y - heightChunks * 3,
                        -80, y - heightChunks * 2]
                },

                { type: 'Q', points: [-75, y, 10, y] }];

            this.body.attr({ d: this.drawPath(body) });
            var leftEye = [
                { type: 'M', points: [-40, y - heightChunks * 2] },
                { type: 'Q', points: [-50, y - heightChunks * 2, -50, y - heightChunks * 2.5] },
                { type: 'Q', points: [-50, y - heightChunks * 3, -40, y - heightChunks * 3] },
                { type: 'Q', points: [-30, y - heightChunks * 3, -30, y - heightChunks * 2.5] },
                { type: 'Q', points: [-30, y - heightChunks * 2, -40, y - heightChunks * 2] }];

            this.eyeLeft.attr({ d: this.drawPath(leftEye) });
            var rightEye = [
                { type: 'M', points: [40, y - heightChunks * 2] },
                { type: 'Q', points: [50, y - heightChunks * 2, 50, y - heightChunks * 2.5] },
                { type: 'Q', points: [50, y - heightChunks * 3, 40, y - heightChunks * 3] },
                { type: 'Q', points: [30, y - heightChunks * 3, 30, y - heightChunks * 2.5] },
                { type: 'Q', points: [30, y - heightChunks * 2, 40, y - heightChunks * 2] }];

            this.eyeRight.attr({ d: this.drawPath(rightEye) });
            var mouth = [
                { type: 'M', points: [0, y - heightChunks * 3] },
                { type: 'Q', points: [20, y - heightChunks * 3, 20, y - heightChunks * 3.5] },
                { type: 'Q', points: [20, y - heightChunks * 4.5, 0, y - heightChunks * 4.5] },
                { type: 'Q', points: [-20, y - heightChunks * 4.5, -20, y - heightChunks * 3.5] },
                { type: 'Q', points: [-20, y - heightChunks * 3, 0, y - heightChunks * 3] }];

            this.mouth.attr({ d: this.drawPath(mouth) });
        }
    };
    return Ghost;
}();
var StageManager = function () {
    function StageManager(svg) {
        this.svg = svg;
        this.ghosts = {};
        this.size = { width: 0, height: 0 };
    }
    StageManager.prototype.init = function () {
        var _this = this;
        window.addEventListener('resize', function () { return _this.onResize(); }, true);
        this.onResize();
        this.tick();
    };
    StageManager.prototype.onResize = function () {
        this.size.width = window.innerWidth;
        this.size.height = window.innerHeight;
        this.svg.attr('width', this.size.width).attr('height', this.size.height);
        // for(let i in this.ghosts)
        // {
        //  this.ghosts[i].updateGuide();
        // }
    };
    StageManager.prototype.addGhost = function () {
        var _this = this;
        var start = { x: this.size.width / 2, y: this.size.height };
        var end = { x: this.size.width / 4 + Math.random() * (this.size.width / 2), y: -300 };
        var ghost = new Ghost(this.svg, start, end, this.onGhostComplete);
        ghost.start(function (ghost) { return _this.removeGhost(ghost); });
        this.ghosts[ghost.id] = ghost;
    };
    StageManager.prototype.removeGhost = function (ghost) {
        delete this.ghosts[ghost.id];
        ghost.remove();
        ghost = null;
    };
    StageManager.prototype.tick = function () {
        var _this = this;
        for (var i in this.ghosts) {
            this.ghosts[i].draw();
        }
        requestAnimationFrame(function () { return _this.tick(); });
    };
    return StageManager;
}();

var stageManager = new StageManager(Snap('svg'));

stageManager.init();

makeGhost();

function makeGhost() {
    stageManager.addGhost();
    setTimeout(makeGhost, Math.random() * 800);
}
