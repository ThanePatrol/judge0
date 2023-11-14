import {EditorView, keymap} from "@codemirror/view"
import {defaultKeymap} from "@codemirror/commands"
import { cpp } from "@codemirror/lang-cpp";
import { python } from "@codemirror/lang-python";
import { java } from "@codemirror/lang-java";
import { basicSetup } from 'codemirror';

let myView = new EditorView({
	doc:
	extensions: [
		basicSetup,

		keymap.of(defaultKeymap)
	],
	parent: document.querySelector("#editor")
})
myView.lineWrapping = true;
myView.dom.style.fontSize = "14px";
myView.theme = "ambiance";


