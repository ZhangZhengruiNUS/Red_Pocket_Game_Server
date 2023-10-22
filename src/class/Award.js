//奖励蜜蜂
import {FlyObj} from './FlyObj.js';
class Award extends FlyObj {
    constructor(width, height, x, y, awardType) {
        super(width, height, x, y);
        this.xSpeed = 5;
        this.ySpeed = 2;
        //1是生命增加，2是火力加强，3是射速,4是威力
        this.awardType = awardType;
    }

    getType() {
        return this.awardType;
    }


    outOfWindow() {
        return this.y > window.innerHeight;
    }


    step() {
        this.x += this.xSpeed;
        this.y += this.ySpeed;
        if (this.x + this.width > window.innerWidth) {
            this.xSpeed = -5;
        }
        if (this.x < 0) {
            this.xSpeed = 5;
        }
    }

}

export default Award;