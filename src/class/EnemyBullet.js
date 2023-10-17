//敌方子弹
import {FlyObj} from './FlyObj.js';

class EnemyBullet extends FlyObj {
    constructor(width, height, x, y, speed, xSpeed) {
        super(width, height, x, y);
        this.speed = speed;
        this.xSpeed = xSpeed;
    }

    step() {
        this.y += this.speed;
        this.x += this.xSpeed;
    }

    outOfWindow() {
        return this.y > window.innerHeight
    }


    //出场之后就闲置
    end() {
        this.x = -800;
        this.y = -800;
        this.speed = 0;
        this.xSpeed = 0;
    }
    isEnd() {
        return this.x === -800;
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
export default EnemyBullet;