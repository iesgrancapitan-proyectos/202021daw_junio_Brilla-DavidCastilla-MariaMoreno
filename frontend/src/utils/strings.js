export function truncate(str, n, useWordBoundary) {
    if (str.length <= n) { return str; }
    const subString = str.substr(0, n - 1); // the original check
    return (useWordBoundary
        ? subString.substr(0, subString.lastIndexOf(" "))
        : subString) + "…";
};

export const merge = (source, target) => {
    return void Object.keys(target).forEach(key => {
        source[key] instanceof Object && target[key] instanceof Object
            ? source[key] instanceof Array && target[key] instanceof Array
                ? void (source[key] = Array.from(new Set(source[key].concat(target[key]))))
                : !(source[key] instanceof Array) && !(target[key] instanceof Array)
                    ? void deepMerge(source[key], target[key])
                    : void (source[key] = target[key])
            : void (source[key] = target[key]);
    }) || source;
}

/** Dispatch event on click outside of node */
export function clickOutside(node) {

  const handleClick = event => {
    if (node && !node.contains(event.target) && !event.defaultPrevented) {
      node.dispatchEvent(
        new CustomEvent('click_outside', node)
      )
    }
  }

	document.addEventListener('click', handleClick, true);

  return {
    destroy() {
      document.removeEventListener('click', handleClick, true);
    }
	}
}
