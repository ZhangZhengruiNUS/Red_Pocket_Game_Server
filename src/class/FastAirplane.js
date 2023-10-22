//超快敌机
import {FlyObj} from './FlyObj.js';
export class FastAirplane extends FlyObj {
    constructor(width, height, x, y) {
        super(width, height, x, y);
        this.speed = 25;
        this.life = 99999;
        this.type = 'FastAirplane'
    }

    subLife() {
        this.life--;
    }

    start() {
        this.life = 99999;
        this.run = true;
    }

    isDie() {
        return this.life <= 0;
    }

    outOfWindow() {
        return this.y > window.innerHeight;
    }

    step() {
        if (this.run) {
            this.y += this.speed
        }

    }
}

export default FastAirplane;
