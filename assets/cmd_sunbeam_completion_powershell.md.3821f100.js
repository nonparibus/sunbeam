import{_ as e,c as s,o,a}from"./app.7b3967da.js";const _=JSON.parse('{"title":"sunbeam completion powershell","description":"","frontmatter":{},"headers":[{"level":2,"title":"Synopsis","slug":"synopsis","link":"#synopsis","children":[]},{"level":2,"title":"Options","slug":"options","link":"#options","children":[]},{"level":2,"title":"See also","slug":"see-also","link":"#see-also","children":[]}],"relativePath":"cmd/sunbeam_completion_powershell.md"}'),n={name:"cmd/sunbeam_completion_powershell.md"},l=a(`<h1 id="sunbeam-completion-powershell" tabindex="-1">sunbeam completion powershell <a class="header-anchor" href="#sunbeam-completion-powershell" aria-hidden="true">#</a></h1><p>Generate the autocompletion script for powershell</p><h2 id="synopsis" tabindex="-1">Synopsis <a class="header-anchor" href="#synopsis" aria-hidden="true">#</a></h2><p>Generate the autocompletion script for powershell.</p><p>To load completions in your current shell session:</p><pre><code>sunbeam completion powershell | Out-String | Invoke-Expression
</code></pre><p>To load completions for every new session, add the output of the above command to your powershell profile.</p><div class="language-"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki material-palenight"><code><span class="line"><span style="color:#A6ACCD;">sunbeam completion powershell [flags]</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span></code></pre></div><h2 id="options" tabindex="-1">Options <a class="header-anchor" href="#options" aria-hidden="true">#</a></h2><div class="language-"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki material-palenight"><code><span class="line"><span style="color:#A6ACCD;">  -h, --help              help for powershell</span></span>
<span class="line"><span style="color:#A6ACCD;">      --no-descriptions   disable completion descriptions</span></span>
<span class="line"><span style="color:#A6ACCD;"></span></span></code></pre></div><h2 id="see-also" tabindex="-1">See also <a class="header-anchor" href="#see-also" aria-hidden="true">#</a></h2><ul><li><a href="./sunbeam_completion">sunbeam completion</a> - Generate the autocompletion script for the specified shell</li></ul>`,12),t=[l];function p(i,r,c,h,d,m){return o(),s("div",null,t)}const b=e(n,[["render",p]]);export{_ as __pageData,b as default};
