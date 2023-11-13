import {EditorView, keymap} from "@codemirror/view"
import {defaultKeymap} from "@codemirror/commands"
import { javascript } from "@codemirror/lang-javascript";
import { python } from "@codemirror/lang-python";
import { basicSetup } from 'codemirror';


let myView = new EditorView({
	doc: 
	extensions: [
		basicSetup,
		keymap.of(defaultKeymap),

	],
	parent: document.querySelector("#editor")
})
myView.lineWrapping = true;
myView.dom.style.fontSize = "14px";
myView.theme = "ambiance";



