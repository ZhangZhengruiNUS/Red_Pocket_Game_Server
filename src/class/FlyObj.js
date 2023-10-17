export class FlyObj {
    constructor(width, height, x = 0, y = 0) {
        this.x = x;
        this.y = y;
        this.width = width;
        this.height = height;
        this.oriX = x;
        this.oriY = y;
        this.run = true
    }

    // 是否出界
    outOfWindow() {

    }

    //飞行物移动一步
    step() {

    }

    //运动
    start() {
        this.run = true;
    }

    //停止
    end() {
        this.run = false;
    }




    //重置到最初位置
    reset(x, y) {
        this.x = x;
        this.y = y;

    }

    shootBy(obj) {
        //如果闲置的飞机就摸不着！！！！！
        if (!this.run) {
            return;
        }

        let x1 = obj.x - this.width / 2;
        let x2 = obj.x + this.width / 2 + obj.width;
        let y1 = obj.y - this.height / 2;
        let y2 = obj.y + this.height / 2 + obj.height;

        let herox = this.x + this.width / 2;
        let heroy = this.y + this.height / 2;

        return herox > x1 && herox < x2 && heroy > y1 && heroy < y2;
    }

}

