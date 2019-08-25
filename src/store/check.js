import axios from "axios"
export default {
  state:{
    check: []
  },
  actions:{
    check(cont, form){ // Действие регистрации
      return new Promise((resolve, reject)=>{
        axios.post('https://delaweb.herokuapp.com/check',{  // https запрос с передачей данных в формате JSON
          email: form.email,
          password: form.password})
            .then(res=>{ //  Если запрос вернул результат
              resolve(res)
            }).catch(res=>{ // Если запрос вернул ошибку
              reject(res)
            })
      })
    }
  }
}
