<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
    </div>
    <button @click="btn">按钮</button>
  </header>

  <main>
    <TheWelcome />
  </main>
</template>

<script>
  
  import { defineComponent } from 'vue'
  // import {axios} from './utils/request'
  // import { useRouter } from 'vue-router'
  import axios from 'axios';
  export default defineComponent({
    name: 'App',
    setup() {
      const router = useRouter()
      const btn = () => {
        router.push('/about')
      }
      return {
        btn
      }
    },
    methods: {
      btn() {
        // axios.post('http://localhost:8080/form_post/',{
        //   message: 'hello',
        //   nick: 'nickname'
        // }).then(res => {
        //   console.log(res)
        // })
        axios.get('http://localhost:8080/welcome?firstname=little&lastname=xi').then(res => {
          console.log(res)
        })
      }
    }
  })


</script>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
