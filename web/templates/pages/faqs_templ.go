// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func addFaqsForm() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!-- Inline form --><div class=\"rounded-2xl border border-gray-200 bg-white dark:border-gray-800 dark:bg-white/[0.03]\" x-show=\"showForm\"><div class=\"p-5 space-y-6 border-t border-gray-100 dark:border-gray-800 sm:p-6\"><form><div class=\"-mx-2.5 flex flex-wrap gap-y-5\"><div class=\"w-full px-2.5\"><label class=\"mb-1.5 block text-sm font-medium text-gray-700 dark:text-gray-400\" for=\"question\">Question</label> <textarea rows=\"3\" id=\"question\" class=\"dark:bg-dark-900 w-full rounded-lg border border-gray-300 bg-transparent px-4 py-2.5 text-sm text-gray-800 shadow-theme-xs placeholder:text-gray-400 focus:border-brand-300 focus:outline-hidden focus:ring-3 focus:ring-brand-500/10 dark:border-gray-700 dark:bg-gray-900 dark:text-white/90 dark:placeholder:text-white/30 dark:focus:border-brand-800\"></textarea></div><div class=\"w-full px-2.5\"><label class=\"mb-1.5 block text-sm font-medium text-gray-700 dark:text-gray-400\" for=\"answer\">Answer</label> <textarea rows=\"3\" id=\"answer\" class=\"dark:bg-dark-900 w-full rounded-lg border border-gray-300 bg-transparent px-4 py-2.5 text-sm text-gray-800 shadow-theme-xs placeholder:text-gray-400 focus:border-brand-300 focus:outline-hidden focus:ring-3 focus:ring-brand-500/10 dark:border-gray-700 dark:bg-gray-900 dark:text-white/90 dark:placeholder:text-white/30 dark:focus:border-brand-800\"></textarea></div></div><div class=\"w-full pt-5 flex justify-end gap-2\"><button type=\"button\" @click=\"showForm = false\" class=\"gap-2 px-4 py-3 text-sm font-medium text-gray-500 rounded-lg bg-white border border-gray-200 hover:bg-gray-50 dark:bg-white/[0.03] dark:text-white/90 dark:border-gray-700 dark:hover:bg-white/[0.03]\">Cancel</button> <button type=\"submit\" class=\"gap-2 px-4 py-3 text-sm font-medium text-white rounded-lg bg-brand-500 hover:bg-brand-600\">Save FAQ</button></div></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func faqsContent() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		return nil
	})
}

func FaqsPage() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<div x-data=\"{ showForm: false }\" class=\"space-y-5 sm:space-y-6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = addFaqsForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<!-- FAQ table --><div class=\"rounded-2xl border border-gray-200 bg-white dark:border-gray-800 dark:bg-white/[0.03]\"><div class=\"px-5 py-4 sm:px-6 sm:py-5 flex items-center justify-between\"><h3 class=\"text-base font-medium text-gray-800 dark:text-white/90\">Predefined FAQs</h3><button @click=\"showForm = !showForm\" class=\"rounded-lg bg-blue-600 px-4 py-2 text-white text-sm hover:bg-blue-700 transition\">Add new FAQ</button></div><div class=\"border-t border-gray-100 p-5 dark:border-gray-800 sm:p-6\"><div class=\"overflow-hidden rounded-xl border border-gray-200 bg-white dark:border-gray-800 dark:bg-white/[0.03]\"><div class=\"max-w-full overflow-x-auto custom-scrollbar\"><table class=\"w-full min-w-[1102px]\"><thead><tr class=\"border-b border-gray-100 dark:border-gray-800\"><th class=\"px-5 py-3 text-left sm:px-6\"><p class=\"font-medium text-gray-500 text-theme-xs dark:text-gray-400\">Question</p></th><th class=\"px-5 py-3 text-left sm:px-6\"><p class=\"font-medium text-gray-500 text-theme-xs dark:text-gray-400\">Answer</p></th></tr></thead> <tbody><tr class=\"border-b border-gray-100 dark:border-gray-800\"><td class=\"px-5 py-4 sm:px-6\"><p class=\"text-gray-600 text-theme-sm dark:text-gray-400\">Website</p></td><td class=\"px-5 py-4 sm:px-6\"><p class=\"text-gray-600 text-theme-sm dark:text-gray-400\">4.5K</p></td></tr></tbody></table></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
