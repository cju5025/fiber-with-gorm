import { Component } from 'react';

export default class AddBookForm extends Component {
    state = {
        title: "",
        author: "",
        rating: 0
    }

    handleSubmit = (event) => {
        event.preventDefault();

        fetch('http://localhost:3000/books', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                title: this.state.title,
                author: this.state.author,
                rating: parseInt(this.state.rating)
            })
        })
    }

    handleChange = (event) => {
        this.setState({
            [event.target.name]: event.target.value
        })
    }

    render() {
        return (
            <div>
                <form id="addBookForm" onSubmit={this.handleSubmit}>
                    <input onChange={this.handleChange} value={this.state.title} name="title" placeholder="title"></input>
                    <input onChange={this.handleChange} value={this.state.author} name="author" placeholder="Author"></input>
                    <input onChange={this.handleChange} value={this.state.rating} name="rating" placeholder="Rating"></input>
                    <input type="Submit"></input>
                </form>
            </div>
        )
    }
}