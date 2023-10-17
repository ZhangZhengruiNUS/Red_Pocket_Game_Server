import {FlyObj} from './FlyObj.js';
class HeroClass extends FlyObj {
    constructor(width, height, x, y) {
        super(width, height, x, y);
        this.life = 100;
        //射击等级
        this.shootLv = 1;
        this.shootSpeed = 1;
        //射击威力
        this.shootPower = 1;
        //1直射 2散射 3乱射
        this.shootType = 1;
    }

    reset(w, h) {
        this.width = w;
        this.height = h;
        this.life = 100;
        //射击等级
        this.shootLv = 1;
        this.shootSpeed = 1;
        //射击威力
        this.shootPower = 1;
        //1直射 2散射 3乱射
        this.shootType = 1;
    }

    addLife(num) {
        this.life += num;
        if (this.life >= 100) {
            this.life = 100;
        }
    }

    subLife(num) {
        this.life -= num;
    }

    supShoot() {
        this.shootLv <= 1 ? 1 : this.shootLv--;
        this.shootSpeed <= 1 ? 1 : this.shootSpeed--;
        this.shootPower <= 1 ? 1 : this.shootPower--;
    }

    addShoot(type) {
        if (type == 2) {
            this.shootLv++;
            if (this.shootLv >= 6) {
                this.shootLv = 6;
            }
        }
        if (type == 3) {
            this.shootSpeed++;
            if (this.shootSpeed >= 3) {
                this.shootSpeed = 3;
            }
        }
        if (type == 4) {
            this.shootPower++;
            if (this.shootPower >= 3) {
                this.shootPower = 3;
            }
        }

    }
    changeShootType() {
        this.shootType = (this.shootType + 1) % 3 + 1;
    }


    isDie() {
        return this.life <= 0;
    }

    moveTo(x, y) {
        this.x = x - this.width / 2;
        this.y = y - this.height / 2;
    }

    outOfWindow() {
        return false
    }

    getAward(type) {
        if (type == 1) {
            this.addLife(10);
        } else {
            this.addShoot(type)
        }
    }

    bulletXspeed(xSpeed) {
        if (this.shootType === 1) {
            return 0;
        }
        if (this.shootType === 2) {
            return xSpeed;
        }
        if (this.shootType === 3) {
            return Math.random() * 2 < 1 ? Math.random() * 20 : -Math.random() * 20
        }
    }


    //hero射击
    shoot() {

        let yStep = 10;

        let arr = [];
        let xStep = this.width / 8;
        switch (this.shootLv) {
            case 2:
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 2,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-5)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 2,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(5)
                })
                break;

            case 3:
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 2,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-5)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 4,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(0)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 6,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(5)
                })
                break;

            case 4:
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 1,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-5)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 3,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-3)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 5,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(3)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 7,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(5)
                })
                break;

            case 5:
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-5)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 2,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-3)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 4,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(0)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 6,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(3)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 8,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(5)
                })
                break;

            case 6:
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * -1,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-5)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 1,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-3)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 3,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(-1)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 5,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(1)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 7,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(3)
                })
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 9,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(5)
                })
                break;
            default:
                arr.push({
                    w: 5,
                    h: 5,
                    x: this.x + xStep * 3,
                    y: this.y - yStep,
                    speed: 10,
                    xSpeed: this.bulletXspeed(0)
                })
                break;

        }


        return arr;
    }


    //被子弹打到或者飞机撞到
    hitBy(obj) {
        let x1 = obj.x - this.width / 2;
        let x2 = obj.x + this.width / 2 + obj.width;
        let y1 = obj.y - this.height / 2;
        let y2 = obj.y + this.height / 2 + obj.height;

        let herox = this.x + this.width / 2;
        let heroy = this.y + this.height / 2;

        return herox > x1 && herox < x2 && heroy > y1 && heroy < y2;
    }

}

export default HeroClass;