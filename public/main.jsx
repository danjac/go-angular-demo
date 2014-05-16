/**
* @jsx React.DOM
*/
'use strict';

var Post = React.createClass({
    
    handleDelete: function () {
        this.props.handleDelete();
        return false;
    },
    render: function () {
        return (
            <li className="post">{this.props.children} <a href="#" onClick={this.handleDelete}>x</a></li>
        )
    }

});

var PostForm = React.createClass({
    handleSubmit: function () {
        var contentNode = this.refs.content.getDOMNode();
        var content = contentNode.value.trim();
        if (!content) {
            return false;
        }
        if (content) {
            this.props.handleSubmit({content: content});
        }
        contentNode.value = "";
        return false;
    },

    render: function () {
        return (
            <form className="PostForm" 
                  onSubmit={this.handleSubmit}>
                <input type="text" 
                       ref="content" 
                       placeholder="post something"/>
                <input type="submit" value="Send"/>
            </form>
        )
    }
});


var PostList = React.createClass({
    
    getInitialState: function () {
        return {data: []};
    },

    handleSubmit: function (post) {
        var _this = this;
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            type: 'POST',
            data: JSON.stringify(post)
        }).success(function (newPost){
            var newData = [newPost].concat(_this.state.data);
            _this.setState({data: newData});
        });
        return false;
    },

    componentWillMount: function () {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            success: function (data) {
                this.setState({data: data}); }.bind(this)
        });
    },

    handleDelete: function(post) {
        var pos = this.state.data.indexOf(post);
        if (pos > -1){
            this.state.data.splice(pos, 1);
            this.setState({data: this.state.data});
            $.ajax({
                url: this.props.url + post.id,
                dataType: "json",
                type: "DELETE"
            });
        }
        return false;
    },

    render: function () {
    
        var handleDelete = this.handleDelete;
        var postNodes = this.state.data.map(function (post) {
            var _handleDelete = function () { handleDelete(post); };
            return <Post key={post.id} 
                          handleDelete={_handleDelete}>{post.content}</Post>
        });

        return (
            <div>
            <PostForm handleSubmit={this.handleSubmit}/>
            <ul className="PostList">
            {postNodes}                            
            </ul>
            </div>
        )

    }
});

React.renderComponent(
    <PostList url="/api/"/>,
    document.getElementById("content")
    );

