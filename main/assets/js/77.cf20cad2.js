(window.webpackJsonp=window.webpackJsonp||[]).push([[77],{640:function(e,t,i){"use strict";i.r(t);var n=i(0),s=Object(n.a)({},(function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[i("h1",{attrs:{id:"state-transitions"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#state-transitions"}},[e._v("#")]),e._v(" State transitions")]),e._v(" "),i("h2",{attrs:{id:"send-fungible-tokens"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#send-fungible-tokens"}},[e._v("#")]),e._v(" Send fungible tokens")]),e._v(" "),i("p",[e._v("A successful fungible token send has two state transitions depending if the transfer is a movement forward or backwards in the token's timeline:")]),e._v(" "),i("ol",[i("li",[e._v("Sender chain is the source chain, "),i("em",[e._v("i.e")]),e._v(" a transfer to any chain other than the one it was previously received from is a movement forwards in the token's timeline. This results in the following state transitions:")])]),e._v(" "),i("ul",[i("li",[e._v("The coins are transferred to an escrow address (i.e locked) on the sender chain.")]),e._v(" "),i("li",[e._v("The coins are transferred to the receiving chain through IBC TAO logic.")])]),e._v(" "),i("ol",{attrs:{start:"2"}},[i("li",[e._v("Sender chain is the sink chain, "),i("em",[e._v("i.e")]),e._v(" the token is sent back to the chain it previously received from. This is a backwards movement in the token's timeline. This results in the following state transitions:")])]),e._v(" "),i("ul",[i("li",[e._v("The coins (vouchers) are burned on the sender chain.")]),e._v(" "),i("li",[e._v("The coins are transferred to the receiving chain through IBC TAO logic.")])]),e._v(" "),i("h2",{attrs:{id:"receive-fungible-tokens"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#receive-fungible-tokens"}},[e._v("#")]),e._v(" Receive fungible tokens")]),e._v(" "),i("p",[e._v("A successful fungible token receive has two state transitions depending if the transfer is a movement forward or backwards in the token's timeline:")]),e._v(" "),i("ol",[i("li",[e._v("Receiver chain is the source chain. This is a backwards movement in the token's timeline. This results in the following state transitions:")])]),e._v(" "),i("ul",[i("li",[e._v("The leftmost port and channel identifier pair is removed from the token denomination prefix.")]),e._v(" "),i("li",[e._v("The tokens are unescrowed and sent to the receiving address.")])]),e._v(" "),i("ol",{attrs:{start:"2"}},[i("li",[e._v("Receiver chain is the sink chain. This is a movement forwards in the token's timeline. This results in the following state transitions:")])]),e._v(" "),i("ul",[i("li",[e._v("Token vouchers are minted by prefixing the destination port and channel identifiers to the trace information.")]),e._v(" "),i("li",[e._v("The receiving chain stores the new trace information in the store (if not set already).")]),e._v(" "),i("li",[e._v("The vouchers are sent to the receiving address.")])])])}),[],!1,null,null,null);t.default=s.exports}}]);