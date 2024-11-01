// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package src

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@2.0.2\"></script><script src=\"https://cdn.tailwindcss.com\"></script></head><body class=\"bg-gradient-to-r from-indigo-500 from-10% via-sky-500 via-30% to-emerald-500 to-90%\"><div class=\"w-screen h-full flex items-center justify-center py-8 bg-gradient-to-r from-indigo-500 from-10% via-sky-500 via-30% to-emerald-500 to-90% \"><div class=\"w-[210mm] bg-white shadow-md rounded-lg border border-gray-300 my-8\"><div class=\"w-full border-b-2\"><p class=\"text-2xl text-center font-mono font-semibold pt-2\">SABIQ AIJAZ</p><p class=\"text-base text-center font-mono font-semibold mb-5\">🏘️ Kashmir 💌 bhatsabiq9@gmail.com 📞 +91-9682161249</p></div><!-- About Section --><section id=\"about\" class=\"p-6\"><h2 class=\"text-3xl font-semibold\">About Me</h2><p class=\"mt-4 text-justify\">I'm a <strong>Full-Stack Software Developer</strong> who is skilled in React and React Native, Go node and express, TypeScript, and Databases (Postgres, MongoDB). Go is my language to build complex and easy-scalable backend servers whereas, React and React Native is easier and faster to implement the users’ interface. I also have experience with creating well-structured web applications with TypeScript and React. And also I know how to use SQL and NOSQL-based databases (Postgres and MongoDB). .</p></section><!-- Skills Section --><section id=\"skills\" class=\"p-6\"><h2 class=\"text-3xl font-semibold\">Skills</h2><div id=\"skill-cloud\" class=\"mt-6\"><div class=\"grid grid-cols-3 gap-4 text-center\"><div class=\"p-4 bg-gray-100 rounded-lg\">Go</div><div class=\"p-4 bg-gray-100 rounded-lg\">Kotlin</div><div class=\"p-4 bg-gray-100 rounded-lg\">Swift</div><div class=\"p-4 bg-gray-100 rounded-lg\">TypeScript</div><div class=\"p-4 bg-gray-100 rounded-lg\">React</div><div class=\"p-4 bg-gray-100 rounded-lg\">React Native</div><div class=\"p-4 bg-gray-100 rounded-lg\">Next.js</div><div class=\"p-4 bg-gray-100 rounded-lg\">Expo</div><div class=\"p-4 bg-gray-100 rounded-lg\">Node</div><div class=\"p-4 bg-gray-100 rounded-lg\">Express</div><div class=\"p-4 bg-gray-100 rounded-lg\">SQL</div><div class=\"p-4 bg-gray-100 rounded-lg\">MongoDB</div></div></div></section><!-- Projects Section --><section id=\"projects\" class=\"p-6\"><h2 class=\"text-3xl font-semibold pb-2\">Projects</h2><div id=\"projects-list\" class=\"flex items-center justify-center\"><button hx-get=\"/projects\" hx-target=\"#projects-list\" hx-swap=\"outerHTML\" class=\"transition-transform transform bg-gradient-to-r from-blue-400 to-blue-600 text-white px-6 py-3 rounded-lg shadow-lg hover:shadow-xl hover:scale-105 focus:outline-none focus:ring-2 focus:ring-blue-300\">Load Projects</button></div></section><!-- Experiece Section --><section id=\"experience\" class=\"p-6\"><h2 class=\"text-3xl font-semibold pb-2\">Experience</h2><div id=\"experience-list\" class=\"flex items-center justify-center\"><button hx-get=\"/experience\" hx-target=\"#experience\" hx-swap=\"outerHTML\" class=\"transition-transform transform bg-gradient-to-r from-blue-400 to-blue-600 text-white px-6 py-3 rounded-lg shadow-lg hover:shadow-xl hover:scale-105 focus:outline-none focus:ring-2 focus:ring-blue-300\">Show Experience</button></div></section><div class=\"w-full bg-gray-800 border-t-2  mt-8 py-6 rounded-b-lg\"><p class=\"text-2xl text-center text-white font-mono font-semibold mb-2\">© SABIQ AIJAZ</p><p class=\"text-base text-center text-gray-300 font-mono\">This website is made with Go, Templ, HTMX, and Tailwind CSS</p></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
