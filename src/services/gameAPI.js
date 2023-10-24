import API from "./API";
class gameAPI{
    static getEquip(useridValue){
        return API().get('/game/equip/'+useridValue,{
            params: {
              userId: useridValue
            }
          })
    }

    static postEquip(userId,equipList){
        return API().post('/game/equip/',
        {
            userId: userId,
            itemId: equipList
        })
    }

    static postCredits(userId,credit,coupon){
        return API().post('/game/end',{
            userId:userId ,
            credit:credit ,
            coupon:coupon
        })
    }

    static  getDiff(){
        return API().get()
    }

    static createPost(data){
        return API('https://jsonplaceholder.typicode.com')
        .post('/posts',data)
    }
}
export default gameAPI;
//     axios.post('',
    //     JSON.stringify({
    //       title:'foo',
    //       body:'bar',
    //       userId:1,
    //     })) .then(response =>{
    //       console.log(response)
    //   })

