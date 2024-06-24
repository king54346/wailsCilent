import HelloWorld from './HelloWorld.vue'

const components = {
    HelloWorld
}

//自动注册全局组件，
export default function setupComponent(app) {
    Object.keys(components).forEach(key => {
        app.component(key, components[key])
    })
}