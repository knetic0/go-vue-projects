const app = Vue.createApp({
    data() {
        return {
            title : "Todo List",
            isShown : [],
            todos : [],
        };
    },
    methods : {
        addTodo() {
            const todo = this.$refs.inputField.value;
            if ( todo.charAt(0) != ' ' && todo.charAt(0) != '' ) {
                this.todos.push(todo)
                this.isShown.push(true)
                this.search = ""
            }
        },
        removeTodo(index) {
            this.todos.splice(index, 1)
        },
        doneTrue(index) {
            if ( this.todos[index].includes("Done!") == false ){
                this.todos[index] += " is Done!"
                console.log(this.todos[index])
                this.isShown[index] = false
            }
        },
        clearTodoList() {
            this.todos.splice(0)
        },
    },
}).mount("#app");