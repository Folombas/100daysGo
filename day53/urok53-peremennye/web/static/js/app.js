(function(){
	function qs(sel, root){return (root||document).querySelector(sel)}
	function qsa(sel, root){return Array.from((root||document).querySelectorAll(sel))}

	function showToast(text){
		let t = qs('.toast');
		if(!t){
			t = document.createElement('div');
			t.className = 'toast';
			document.body.appendChild(t);
		}
		t.textContent = text;
		t.classList.add('show');
		setTimeout(()=>t.classList.remove('show'), 1600);
	}

	document.addEventListener('click', function(e){
		const btn = e.target.closest('button.btn');
		if(!btn) return;
		const action = btn.getAttribute('data-action');
		const targetId = btn.getAttribute('data-target');
		const el = targetId ? qs('#'+CSS.escape(targetId)) : null;
		if(action === 'toggle' && el){
			el.classList.toggle('hidden');
		}
		if(action === 'copy' && el){
			const text = el.innerText;
			navigator.clipboard.writeText(text).then(function(){
				showToast('Скопировано в буфер обмена');
			}).catch(function(){
				showToast('Не удалось скопировать');
			});
		}
	});
})();
