<template>
    <div>
      <h1>My Book List</h1>
      <button @click="showAddForm = true">Add New Book</button>
  
      <div v-if="showAddForm">
        <h2>Add Book</h2>
        <form @submit.prevent="addBook">
          <input v-model="newBook.title" placeholder="Title" required />
          <input v-model="newBook.author" placeholder="Author" required />
          <select v-model="newBook.status">
            <option value="to-read">To Read</option>
            <option value="reading">Reading</option>
            <option value="completed">Completed</option>
          </select>
          <button type="submit">Add</button>
          <button @click="showAddForm = false">Cancel</button>
        </form>
      </div>
  
      <ul>
        <li v-for="book in books" :key="book.id">
          <strong>{{ book.title }}</strong> by {{ book.author }} - {{ book.status }}
          <button @click="deleteBook(book.id)">Delete</button>
          <button @click="updateBookStatus(book)">Update Status</button>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        books: [],
        newBook: {
          title: '',
          author: '',
          status: 'to-read',
        },
        showAddForm: false,
      };
    },
    methods: {
      fetchBooks() {
        axios.get('http://localhost:8080/books')
          .then(response => {
            this.books = response.data;
          })
          .catch(error => {
            console.error(error);
          });
      },
      addBook() {
        axios.post('http://localhost:8080/books', this.newBook)
          .then(() => {
            this.fetchBooks();
            this.showAddForm = false;
            this.newBook = { title: '', author: '', status: 'to-read' };
            console.log('added')
          })
          .catch(error => {
            console.error(error);
          });
      },
      deleteBook(id) {
        axios.delete(`http://localhost:8080/books/${id}`)
          .then(() => {
            this.fetchBooks();
          })
          .catch(error => {
            console.error(error);
          });
      },
      updateBookStatus(book) {
        const statuses = ['to-read', 'reading', 'completed'];
        const currentIndex = statuses.indexOf(book.status);
        book.status = statuses[(currentIndex + 1) % statuses.length];
        axios.put(`http://localhost:8080/books/${book.id}`, book)
          .then(() => {
            this.fetchBooks();
          })
          .catch(error => {
            console.error(error);
          });
      },
    },
    mounted() {
      this.fetchBooks();
    },
  };
  </script>
  
  <style scoped>
  /* Add some basic styling */
  </style>
  