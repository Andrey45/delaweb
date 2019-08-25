// Чтобы все выйлы собирались с приставой пути ../assets нужно для сервера
module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? '../assets'
    : ''
}
