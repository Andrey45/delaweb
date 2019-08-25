import axios from "axios"
export default {
  state:{
    login: false,
    user: ''
  },
  actions:{
    login(cont, form){ // Действие авторизации
      return new Promise((resolve, reject) => {
        axios.post('https://delaweb.herokuapp.com/login', { // https запрос с передачей данных в формате JSON
          email: form.email,
          password: form.password
        }).then(res =>{ //  Если запрос вернул результат
          cont.commit('login', res.data)
          resolve(res)
        }).catch(err=>{ // Если запрос вернул ошибку
          reject(err)
        })
      })
    },
    updeteUser(cont, form){ // Действие изменение данных профиля
      return new  Promise((resolve, reject)=>{
        axios.put('https://delaweb.herokuapp.com/userUpt', { // https запрос с передачей данных в формате JSON
          email: form.email,
          username: form.username,
          country: form.country,
          city: form.city
        }).then(res=>{ //  Если запрос вернул результат
          cont.commit('updateUser', res.data)
          resolve(res)
        }).catch(err=>{ // Если запрос вернул ошибку
          reject(err)
        })
      })
    }
  },
  mutations:{
    login(state, logo){ // Мутация авторизации
      state.login = !state.login
      state.user = logo
      localStorage.setItem('userDeLaWeb', JSON.stringify(logo))
    },
    Logout(state){ // Мутация выходи из профиля
      state.user = []
      state.login = !state.login
      localStorage.removeItem('userDeLaWeb')
    },
    updateUser(state, user){ // Мутация изменения данных профиля
      state.user = user
      localStorage.setItem('userDeLaWeb', JSON.stringify(user))
    }
  },
  getters:{
    Login(state){ // Геттер состояния аворизации пользователя
      if(state.login === false && localStorage.getItem('userDeLaWeb') === null){ // Если state login отрицателен и в localStorage нет данного ключа то возврящает отрицание
        return state.login
      } else if (state.login === true && localStorage.getItem('userDeLaWeb') === null) { // Если state login положителен и в localStorage нет данного ключа то возврящает отрицание
        return false
      } else {  // Иначе возвращает положительный результат
        return true
      }
    },
    UserData(state){ // Геттер возврящает данные о пользователи
      if(state.user !== ''){
        return state.user
      }else {
        if(localStorage.getItem('userDeLaWeb') !== null) { // Если ключь существует значит пользователь был авторизован в прошлой сесии
          state.user = JSON.parse(localStorage.getItem('userDeLaWeb'))
          return state.user
        } else {
          return  false
        }
      }
    }
  }
}
