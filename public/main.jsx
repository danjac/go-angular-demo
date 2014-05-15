/**
* @jsx React.DOM
*/
'use strict';

var Tweet = React.createClass({
    
    handleDelete: function () {
        this.props.handleDelete();
        return false;
    },
    render: function () {
        return (
            <li className="Tweet">{this.props.children} <a href="#" onClick={this.handleDelete}>x</a></li>
        )
    }

});

var TweetForm = React.createClass({
    handleSubmit: function () {
        var contentNode = this.refs.content.getDOMNode();
        var content = contentNode.value.trim();
        if (content) {
            this.props.handleSubmit({content: content});
        }
        contentNode.value = "";
    },

    render: function () {
        return (
            <form className="TweetForm" 
                  onSubmit={this.handleSubmit}>
                <input type="text" 
                       ref="content" 
                       placeholder="tweet something"/>
                <input type="submit" value="Send"/>
            </form>
        )
    }
});


var TweetList = React.createClass({
    
    getInitialState: function () {
        return {data: []};
    },

    handleSubmit: function (tweet) {
        this.state.data.splice(0, 0, tweet);
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            type: 'POST',
            data: JSON.stringify(tweet)
        });
    },

    componentWillMount: function () {
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            success: function (data) {
                this.setState({data: data}); }.bind(this)
        });
    },

    handleDelete: function(tweet) {
        var pos = this.state.data.indexOf(tweet);
        if (pos > -1){
            this.state.data.splice(pos, 1);
            this.setState({data: this.state.data});
            $.ajax({
                url: this.props.url + tweet.id,
                dataType: "json",
                type: "DELETE"
            });
        }
    },

    render: function () {
    
        var handleDelete = this.handleDelete;
        var tweetNodes = this.state.data.map(function (tweet) {
            var _handleDelete = function () { handleDelete(tweet); };
            return <Tweet key={tweet.id} 
                          handleDelete={_handleDelete}>{tweet.content}</Tweet>
        });

        return (
            <div>
            <TweetForm handleSubmit={this.handleSubmit}/>
            <ul className="TweetList">
            {tweetNodes}                            
            </ul>
            </div>
        )

    }
});

React.renderComponent(
    <TweetList url="/api/"/>,
    document.getElementById("content")
    );

