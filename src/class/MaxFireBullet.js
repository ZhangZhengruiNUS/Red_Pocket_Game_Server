//绝杀子弹
import {FlyObj} from './FlyObj.js';
export class MaxFireBullet extends FlyObj {
    constructor(width, height, x, y, speed, life = 0) {
        super(width, height, x, y);
        this.speed = speed;
        this.life = life;
    }
    outOfWindow() {
        return this.y < 0;
    }
    step() {
        this.y -= this.speed;
    }
    end() {
        this.x = -1200;
        this.y = -1200;
        this.speed = 0;
    }
    subLife(num) {
        this.life -= num;
    }
    isDie() {
        return this.life <= 0
    }
    isEnd() {
        return this.x === -1200;
    }
    reset(w, h, x, y, speed, life) {
        this.width = w;
        this.height = h;
        this.x = x;
        this.y = y;
        this.speed = speed;
        this.life = life;
    }

}

export default MaxFireBullet;
