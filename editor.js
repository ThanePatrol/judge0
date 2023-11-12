import {basicSetup, EditorView} from "codemirror"
import {EditorState, Compartment} from "@codemirror/state"
import {python} from "@codemirror/lang-python"

let language = new Compartment, tabSize = new Compartment

let state = EditorState.create({
  extensions: [
    basicSetup,
    language.of(python()),
    tabSize.of(EditorState.tabSize.of(8))
  ]
})

let view = new EditorView({
  state,
  parent: document.body
})
