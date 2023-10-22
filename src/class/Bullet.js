//子弹
import {FlyObj} from './FlyObj.js';
class Bullet extends FlyObj {
    constructor(width, height, x, y, speed, xSpeed) {
        super(width, height, x, y);
        this.speed = speed;
        this.xSpeed = xSpeed
    }

    step() {
        this.y -= this.speed;
        this.x += this.xSpeed;
    }

    outOfWindow() {
        return this.y < 0;
    }


    //出场之后就闲置
    end() {
        this.x = -400;
        this.y = -400;
        this.speed = 0;
        this.xSpeed = 0;
    }
    isEnd() {
        return this.x === -400;
    }

    reset(w, h, x, y, speed, xSpeed) {
        this.width = w;
        this.height = h;
        this.x = x;
        this.y = y;
        this.speed = speed;
        this.xSpeed = xSpeed;
    }
}

export default Bullet;