webpackHotUpdate_N_E("pages/index",{

/***/ "./src/components/InputField.js":
/*!**************************************!*\
  !*** ./src/components/InputField.js ***!
  \**************************************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* WEBPACK VAR INJECTION */(function(module) {/* harmony import */ var _Users_jcarrill_go_src_github_com_rillo_carrillo_restaurant_web_node_modules_next_node_modules_babel_runtime_helpers_esm_defineProperty__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./node_modules/next/node_modules/@babel/runtime/helpers/esm/defineProperty */ \"./node_modules/next/node_modules/@babel/runtime/helpers/esm/defineProperty.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var _Users_jcarrill_go_src_github_com_rillo_carrillo_restaurant_web_node_modules_next_node_modules_babel_runtime_helpers_esm_slicedToArray__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./node_modules/next/node_modules/@babel/runtime/helpers/esm/slicedToArray */ \"./node_modules/next/node_modules/@babel/runtime/helpers/esm/slicedToArray.js\");\n/* harmony import */ var _Users_jcarrill_go_src_github_com_rillo_carrillo_restaurant_web_node_modules_next_node_modules_babel_runtime_helpers_esm_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./node_modules/next/node_modules/@babel/runtime/helpers/esm/objectWithoutProperties */ \"./node_modules/next/node_modules/@babel/runtime/helpers/esm/objectWithoutProperties.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_4__);\n/* harmony import */ var prop_types__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! prop-types */ \"./node_modules/prop-types/index.js\");\n/* harmony import */ var prop_types__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(prop_types__WEBPACK_IMPORTED_MODULE_5__);\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @chakra-ui/react */ \"./node_modules/@chakra-ui/react/dist/esm/index.js\");\n/* harmony import */ var formik__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! formik */ \"./node_modules/formik/dist/formik.esm.js\");\n\n\n\n\n\nvar _jsxFileName = \"/Users/jcarrill/go/src/github.com/rillo-carrillo/restaurant/web/src/components/InputField.js\",\n    _this = undefined,\n    _s = $RefreshSig$();\n\nfunction ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }\n\nfunction _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { Object(_Users_jcarrill_go_src_github_com_rillo_carrillo_restaurant_web_node_modules_next_node_modules_babel_runtime_helpers_esm_defineProperty__WEBPACK_IMPORTED_MODULE_0__[\"default\"])(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }\n\n\n\n\n\n\nvar InputField = function InputField(_ref) {\n  _s();\n\n  var id = _ref.id,\n      label = _ref.label,\n      err = _ref.err,\n      props = Object(_Users_jcarrill_go_src_github_com_rillo_carrillo_restaurant_web_node_modules_next_node_modules_babel_runtime_helpers_esm_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3__[\"default\"])(_ref, [\"id\", \"label\", \"err\"]);\n\n  var _useField = Object(formik__WEBPACK_IMPORTED_MODULE_7__[\"useField\"])(props),\n      _useField2 = Object(_Users_jcarrill_go_src_github_com_rillo_carrillo_restaurant_web_node_modules_next_node_modules_babel_runtime_helpers_esm_slicedToArray__WEBPACK_IMPORTED_MODULE_2__[\"default\"])(_useField, 2),\n      field = _useField2[0],\n      error = _useField2[1].error;\n\n  var C = _chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__[\"Input\"];\n  console.log(error !== undefined);\n  console.log(error);\n  return /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__[\"jsxDEV\"])(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__[\"Box\"], {\n    mt: 4,\n    children: /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__[\"jsxDEV\"])(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__[\"FormControl\"], {\n      id: id,\n      isInvalid: !!error,\n      children: [/*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__[\"jsxDEV\"])(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__[\"FormLabel\"], {\n        htmlFor: field.name,\n        children: label\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 19,\n        columnNumber: 9\n      }, _this), /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__[\"jsxDEV\"])(C, _objectSpread(_objectSpread(_objectSpread({}, field), props), {}, {\n        id: field.name\n      }), void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 20,\n        columnNumber: 9\n      }, _this), error !== undefined ? /*#__PURE__*/Object(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_1__[\"jsxDEV\"])(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__[\"FormErrorMessage\"], {\n        children: error\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 22,\n        columnNumber: 11\n      }, _this) : null]\n    }, void 0, true, {\n      fileName: _jsxFileName,\n      lineNumber: 18,\n      columnNumber: 7\n    }, _this)\n  }, void 0, false, {\n    fileName: _jsxFileName,\n    lineNumber: 17,\n    columnNumber: 5\n  }, _this);\n};\n\n_s(InputField, \"DpPuTb9nPdu6NyyZR63Xu5KSrwc=\", false, function () {\n  return [formik__WEBPACK_IMPORTED_MODULE_7__[\"useField\"]];\n});\n\n_c = InputField;\nInputField.propTypes = {\n  id: prop_types__WEBPACK_IMPORTED_MODULE_5___default.a.string,\n  label: prop_types__WEBPACK_IMPORTED_MODULE_5___default.a.string,\n  err: prop_types__WEBPACK_IMPORTED_MODULE_5___default.a.bool\n};\n/* harmony default export */ __webpack_exports__[\"default\"] = (InputField);\n\nvar _c;\n\n$RefreshReg$(_c, \"InputField\");\n\n;\n    var _a, _b;\n    // Legacy CSS implementations will `eval` browser code in a Node.js context\n    // to extract CSS. For backwards compatibility, we need to check we're in a\n    // browser context before continuing.\n    if (typeof self !== 'undefined' &&\n        // AMP / No-JS mode does not inject these helpers:\n        '$RefreshHelpers$' in self) {\n        var currentExports = module.__proto__.exports;\n        var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n        // This cannot happen in MainTemplate because the exports mismatch between\n        // templating and execution.\n        self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.i);\n        // A module can be accepted automatically based on its exports, e.g. when\n        // it is a Refresh Boundary.\n        if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n            // Save the previous exports on update so we can compare the boundary\n            // signatures.\n            module.hot.dispose(function (data) {\n                data.prevExports = currentExports;\n            });\n            // Unconditionally accept an update to this module, we'll check if it's\n            // still a Refresh Boundary later.\n            module.hot.accept();\n            // This field is set when the previous version of this module was a\n            // Refresh Boundary, letting us know we need to check for invalidation or\n            // enqueue an update.\n            if (prevExports !== null) {\n                // A boundary can become ineligible if its exports are incompatible\n                // with the previous exports.\n                //\n                // For example, if you add/remove/change exports, we'll want to\n                // re-execute the importing modules, and force those components to\n                // re-render. Similarly, if you convert a class component to a\n                // function, we want to invalidate the boundary.\n                if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                    module.hot.invalidate();\n                }\n                else {\n                    self.$RefreshHelpers$.scheduleUpdate();\n                }\n            }\n        }\n        else {\n            // Since we just executed the code for the module, it's possible that the\n            // new exports made it ineligible for being a boundary.\n            // We only care about the case when we were _previously_ a boundary,\n            // because we already accepted this update (accidental side effect).\n            var isNoLongerABoundary = prevExports !== null;\n            if (isNoLongerABoundary) {\n                module.hot.invalidate();\n            }\n        }\n    }\n\n/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(/*! ./../../node_modules/next/dist/compiled/webpack/harmony-module.js */ \"./node_modules/next/dist/compiled/webpack/harmony-module.js\")(module)))//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vc3JjL2NvbXBvbmVudHMvSW5wdXRGaWVsZC5qcz8wNWM3Il0sIm5hbWVzIjpbIklucHV0RmllbGQiLCJpZCIsImxhYmVsIiwiZXJyIiwicHJvcHMiLCJ1c2VGaWVsZCIsImZpZWxkIiwiZXJyb3IiLCJDIiwiSW5wdXQiLCJjb25zb2xlIiwibG9nIiwidW5kZWZpbmVkIiwibmFtZSIsInByb3BUeXBlcyIsIlByb3BUeXBlcyIsInN0cmluZyIsImJvb2wiXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUNBO0FBQ0E7QUFPQTs7QUFDQSxJQUFNQSxVQUFVLEdBQUcsU0FBYkEsVUFBYSxPQUFrQztBQUFBOztBQUFBLE1BQS9CQyxFQUErQixRQUEvQkEsRUFBK0I7QUFBQSxNQUEzQkMsS0FBMkIsUUFBM0JBLEtBQTJCO0FBQUEsTUFBcEJDLEdBQW9CLFFBQXBCQSxHQUFvQjtBQUFBLE1BQVpDLEtBQVk7O0FBQUEsa0JBQ3hCQyx1REFBUSxDQUFDRCxLQUFELENBRGdCO0FBQUE7QUFBQSxNQUM1Q0UsS0FENEM7QUFBQSxNQUNuQ0MsS0FEbUMsaUJBQ25DQSxLQURtQzs7QUFFbkQsTUFBSUMsQ0FBQyxHQUFHQyxzREFBUjtBQUNBQyxTQUFPLENBQUNDLEdBQVIsQ0FBWUosS0FBSyxLQUFLSyxTQUF0QjtBQUNBRixTQUFPLENBQUNDLEdBQVIsQ0FBWUosS0FBWjtBQUNBLHNCQUNFLHFFQUFDLG9EQUFEO0FBQUssTUFBRSxFQUFFLENBQVQ7QUFBQSwyQkFDRSxxRUFBQyw0REFBRDtBQUFhLFFBQUUsRUFBRU4sRUFBakI7QUFBcUIsZUFBUyxFQUFFLENBQUMsQ0FBQ00sS0FBbEM7QUFBQSw4QkFDRSxxRUFBQywwREFBRDtBQUFXLGVBQU8sRUFBRUQsS0FBSyxDQUFDTyxJQUExQjtBQUFBLGtCQUFpQ1g7QUFBakM7QUFBQTtBQUFBO0FBQUE7QUFBQSxlQURGLGVBRUUscUVBQUMsQ0FBRCxnREFBT0ksS0FBUCxHQUFrQkYsS0FBbEI7QUFBeUIsVUFBRSxFQUFFRSxLQUFLLENBQUNPO0FBQW5DO0FBQUE7QUFBQTtBQUFBO0FBQUEsZUFGRixFQUdHTixLQUFLLEtBQUtLLFNBQVYsZ0JBQ0MscUVBQUMsaUVBQUQ7QUFBQSxrQkFBbUJMO0FBQW5CO0FBQUE7QUFBQTtBQUFBO0FBQUEsZUFERCxHQUVHLElBTE47QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBREY7QUFBQTtBQUFBO0FBQUE7QUFBQSxXQURGO0FBV0QsQ0FoQkQ7O0dBQU1QLFU7VUFDdUJLLCtDOzs7S0FEdkJMLFU7QUFrQk5BLFVBQVUsQ0FBQ2MsU0FBWCxHQUF1QjtBQUNyQmIsSUFBRSxFQUFFYyxpREFBUyxDQUFDQyxNQURPO0FBRXJCZCxPQUFLLEVBQUVhLGlEQUFTLENBQUNDLE1BRkk7QUFHckJiLEtBQUcsRUFBRVksaURBQVMsQ0FBQ0U7QUFITSxDQUF2QjtBQU1lakIseUVBQWYiLCJmaWxlIjoiLi9zcmMvY29tcG9uZW50cy9JbnB1dEZpZWxkLmpzLmpzIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IFJlYWN0IGZyb20gXCJyZWFjdFwiO1xuaW1wb3J0IFByb3BUeXBlcyBmcm9tIFwicHJvcC10eXBlc1wiO1xuaW1wb3J0IHtcbiAgRm9ybUNvbnRyb2wsXG4gIEZvcm1MYWJlbCxcbiAgSW5wdXQsXG4gIEJveCxcbiAgRm9ybUVycm9yTWVzc2FnZSxcbn0gZnJvbSBcIkBjaGFrcmEtdWkvcmVhY3RcIjtcbmltcG9ydCB7IHVzZUZpZWxkIH0gZnJvbSBcImZvcm1pa1wiO1xuY29uc3QgSW5wdXRGaWVsZCA9ICh7IGlkLCBsYWJlbCwgZXJyLCAuLi5wcm9wcyB9KSA9PiB7XG4gIGNvbnN0IFtmaWVsZCwgeyBlcnJvciB9XSA9IHVzZUZpZWxkKHByb3BzKTtcbiAgbGV0IEMgPSBJbnB1dDtcbiAgY29uc29sZS5sb2coZXJyb3IgIT09IHVuZGVmaW5lZCk7XG4gIGNvbnNvbGUubG9nKGVycm9yKTtcbiAgcmV0dXJuIChcbiAgICA8Qm94IG10PXs0fT5cbiAgICAgIDxGb3JtQ29udHJvbCBpZD17aWR9IGlzSW52YWxpZD17ISFlcnJvcn0+XG4gICAgICAgIDxGb3JtTGFiZWwgaHRtbEZvcj17ZmllbGQubmFtZX0+e2xhYmVsfTwvRm9ybUxhYmVsPlxuICAgICAgICA8QyB7Li4uZmllbGR9IHsuLi5wcm9wc30gaWQ9e2ZpZWxkLm5hbWV9IC8+XG4gICAgICAgIHtlcnJvciAhPT0gdW5kZWZpbmVkID8gKFxuICAgICAgICAgIDxGb3JtRXJyb3JNZXNzYWdlPntlcnJvcn08L0Zvcm1FcnJvck1lc3NhZ2U+XG4gICAgICAgICkgOiBudWxsfVxuICAgICAgPC9Gb3JtQ29udHJvbD5cbiAgICA8L0JveD5cbiAgKTtcbn07XG5cbklucHV0RmllbGQucHJvcFR5cGVzID0ge1xuICBpZDogUHJvcFR5cGVzLnN0cmluZyxcbiAgbGFiZWw6IFByb3BUeXBlcy5zdHJpbmcsXG4gIGVycjogUHJvcFR5cGVzLmJvb2wsXG59O1xuXG5leHBvcnQgZGVmYXVsdCBJbnB1dEZpZWxkO1xuIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./src/components/InputField.js\n");

/***/ })

})