package views

const Styles = `
	html, body {
		height: 100%;
	}
	#left {
		display: flex;
		flex-flow: column;
		height: 100%;
	}
	.menu {
		min-height: 56px;
	}
	.editor, .empty-panel {
		flex: 1;
		width: 100%;
	}
	.empty-panel {
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.split {
		height: 100%;
		width: 100%;
	}
	.gutter {
		height: 100%;
		background-color: #eee;
		background-repeat: no-repeat;
		background-position: 50%;
	}
	.gutter.gutter-horizontal {
		cursor: col-resize;
		background-image:  url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAeCAYAAADkftS9AAAAIklEQVQoU2M4c+bMfxAGAgYYmwGrIIiDjrELjpo5aiZeMwF+yNnOs5KSvgAAAABJRU5ErkJggg==')
	}
	.gutter.gutter-vertical {
		cursor: row-resize;
		background-image:  url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAFAQMAAABo7865AAAABlBMVEVHcEzMzMzyAv2sAAAAAXRSTlMAQObYZgAAABBJREFUeF5jOAMEEAIEEFwAn3kMwcB6I2AAAAAASUVORK5CYII=')
	}
	.split {
		-webkit-box-sizing: border-box;
		-moz-box-sizing: border-box;
		box-sizing: border-box;
	}
	.split, .gutter.gutter-horizontal {
		float: left;
	}
	.preview {
		border: 0;
		height: 100%;
		width: 100%;
	}
	#console-holder {
		overflow: auto;
	}
	#console {
		padding:5px;
	}
	.octicon {
		display: inline-block;
		vertical-align: text-top;
		fill: currentColor;
	}
	#help-modal table { 
		clear: both;
	}
	#help-modal img { 
		margin-left: 20px;
		margin-bottom: 30px;
	}
	#help-modal h2 { 
		padding-bottom: 0.3em;
    	font-size: 1.5em;
    	border-bottom: 1px solid #eaecef;
		margin-top: 24px;
		margin-bottom: 16px;
		font-weight: 600;
		line-height: 1.25;
	}
	#help-modal h4 {
	    font-size: 1em;
		margin-top: 24px;
    	margin-bottom: 16px;
    	font-weight: 600;
    	line-height: 1.25;
	}
	#help-modal .modal-lg {
		max-width: 700px;
	}
	#help-modal a {
		color: #0366d6;
	}
	#help-modal code {
		padding: 0.2em 0.4em;
		margin: 0;
		font-size: 85%;
		background-color: rgba(27,31,35,0.05);
		border-radius: 3px;
		color: #24292e;
	}
`
