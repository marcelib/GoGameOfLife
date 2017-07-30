<script>
  require('es6-promise').polyfill();
  require('isomorphic-fetch');

  export default {
    name: 'hello',
    data() {
      return {
        title: 'Title',
        description: '',
      };
    },
    created() {
      fetch('http://localhost:8080/index')
        .then((response) => {
          if (response.status >= 400) {
            throw new Error('Bad response from server');
          }
          return response.json();
        })
        .then((body) => {
          this.title = body.title;
          this.description = body.description;
        });
    },
  };
</script>

<template>
  <div class="hello">
    <h1>{{ title }}</h1>
    <h2>{{ description }}</h2>
  </div>
</template>


<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
