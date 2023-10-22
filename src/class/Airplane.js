//普通敌机
import {FlyObj} from './FlyObj.js';
class Airplane extends FlyObj {
    constructor(width, height, x, y) {
        super(width, height, x, y);
        this.speed = 3;
        this.life = 3;
        this.score = 3;
        this.type = 'Airplane'
    }

    getScore() {
        return this.score;
    }

    subLife(num) {
        this.life -= num;
    }

    start() {
        this.life = 3;
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

export default Airplane;