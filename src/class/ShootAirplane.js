//发射敌机
import {FlyObj} from './FlyObj.js';
export class ShootAirplane extends FlyObj {
    constructor(width, height, x, y) {
        super(width, height, x, y);
        this.speed = 3;
        this.life = 6;
        this.score = 6;
        this.canShoot = true;
        this.type = 'ShootAirplane'
    }

    getScore() {
        return this.score;
    }

    subLife() {
        this.life--;
    }

    isDie() {
        return this.life <= 0;
    }

    start() {
        this.life = 5;
        this.run = true;
    }

    outOfWindow() {
        return this.y > window.innerHeight;
    }

    shoot() {

        let xStep = this.width / 3;
        let yStep = 20;

        let arr = [];
        // arr.push(new EnemyBullet(this.x + xStep, this.y + yStep, 5, 5, 20))
        let obj = {
            width: 8,
            height: 8,
            x: this.x + xStep + 7,
            y: this.y + yStep,
            speed: 15,
            xSpeed: this.x < (window.innerWidth / 2) ? 3 * Math.random() : -3 * Math.random()
        }

        return obj;
    }

    step() {
        if (this.run) {
            this.y += this.speed
        }

    }
}

export default ShootAirplane;