module.exports = {

"[externals]/next/dist/compiled/next-server/app-page.runtime.dev.js [external] (next/dist/compiled/next-server/app-page.runtime.dev.js, cjs)": (function(__turbopack_context__) {

var { g: global, __dirname, m: module, e: exports } = __turbopack_context__;
{
const mod = __turbopack_context__.x("next/dist/compiled/next-server/app-page.runtime.dev.js", () => require("next/dist/compiled/next-server/app-page.runtime.dev.js"));

module.exports = mod;
}}),
"[project]/src/utils/api.ts [app-ssr] (ecmascript)": ((__turbopack_context__) => {
"use strict";

var { g: global, __dirname } = __turbopack_context__;
{
__turbopack_context__.s({
    "createArticle": (()=>createArticle),
    "createUser": (()=>createUser),
    "deleteArticle": (()=>deleteArticle),
    "getArticle": (()=>getArticle),
    "getArticles": (()=>getArticles),
    "getUsers": (()=>getUsers),
    "login": (()=>login),
    "updateArticle": (()=>updateArticle)
});
const API_URL = ("TURBOPACK compile-time truthy", 1) ? ("TURBOPACK compile-time value", "http://backend:8080/api/v1") || "http://backend:8080/api/v1" // URL pour SSR
 : ("TURBOPACK unreachable", undefined); // URL pour le navigateur
const AUTH_URL = ("TURBOPACK compile-time truthy", 1) ? ("TURBOPACK compile-time value", "http://auth-service:8081/api/v1") || "http://auth-service:8081/api/v1" // URL pour SSR
 : ("TURBOPACK unreachable", undefined); // URL pour le navigateur
// Ajout d'une fonction utilitaire pour vérifier l'authentification
function checkAuth() {
    if ("TURBOPACK compile-time falsy", 0) {
        "TURBOPACK unreachable";
    }
}
async function getArticles() {
    try {
        // Vérifier l'authentification avant de charger les articles
        checkAuth();
        console.log('Environment:', ("TURBOPACK compile-time truthy", 1) ? 'server' : ("TURBOPACK unreachable", undefined));
        console.log('Fetching from:', `${API_URL}/articles`);
        const response = await fetch(`${API_URL}/articles`, {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            mode: 'cors',
            cache: 'no-cache'
        });
        if (!response.ok) {
            console.error('Response not OK:', response.status, response.statusText);
            const errorData = await response.json().catch(()=>null);
            throw new Error(errorData?.message || "Failed to fetch articles");
        }
        const data = await response.json();
        console.log('Received data:', data);
        return data;
    } catch (error) {
        console.error("Error fetching articles:", error);
        throw error;
    }
}
async function getArticle(id) {
    const response = await fetch(`${API_URL}/articles/${id}`);
    if (!response.ok) throw new Error("Failed to fetch article");
    return response.json();
}
async function createArticle(article) {
    try {
        checkAuth();
        const response = await fetch(`${API_URL}/articles`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify(article),
            mode: "cors"
        });
        if (!response.ok) {
            const errorData = await response.json().catch(()=>null);
            throw new Error(errorData?.message || "Failed to create article");
        }
        return response.json();
    } catch (error) {
        console.error("Error creating article:", error);
        throw error;
    }
}
async function updateArticle(id, article) {
    try {
        checkAuth();
        console.log('Updating article:', id, article);
        const response = await fetch(`${API_URL}/articles/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            mode: "cors",
            body: JSON.stringify(article)
        });
        if (!response.ok) {
            const errorData = await response.json().catch(()=>null);
            console.error('Update error response:', errorData);
            throw new Error(errorData?.message || "Failed to update article");
        }
        return response.json();
    } catch (error) {
        console.error("Error updating article:", error);
        throw error;
    }
}
async function deleteArticle(id) {
    try {
        checkAuth();
        console.log('Deleting article:', id);
        const response = await fetch(`${API_URL}/articles/${id}`, {
            method: "DELETE",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            mode: "cors"
        });
        if (!response.ok) {
            const errorData = await response.json().catch(()=>null);
            console.error('Delete error response:', errorData);
            throw new Error(errorData?.message || "Failed to delete article");
        }
    } catch (error) {
        console.error("Error deleting article:", error);
        throw error;
    }
}
async function login(username, password) {
    try {
        console.log('Logging in with:', username);
        console.log('Auth URL:', AUTH_URL);
        const response = await fetch(`${AUTH_URL}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            mode: "cors",
            body: JSON.stringify({
                username,
                password
            })
        });
        if (!response.ok) {
            console.error('Login failed:', response.status, response.statusText);
            const errorData = await response.json().catch(()=>null);
            throw new Error(errorData?.error || "Login failed");
        }
        const data = await response.json();
        console.log('Login response:', data);
        return data;
    } catch (error) {
        console.error("Error during login:", error);
        throw error;
    }
}
async function getUsers() {
    try {
        const response = await fetch(`${AUTH_URL}/users`, {
            method: "GET",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            mode: "cors"
        });
        if (!response.ok) {
            const errorData = await response.json().catch(()=>null);
            throw new Error(errorData?.error || "Failed to fetch users");
        }
        return response.json();
    } catch (error) {
        console.error("Error fetching users:", error);
        throw error;
    }
}
async function createUser(userData) {
    try {
        console.log('Creating user with data:', {
            username: userData.username,
            email: userData.email,
            passwordProvided: !!userData.password,
            passwordLength: userData.password?.length || 0
        });
        // Vérification supplémentaire côté client
        if (!userData.username || !userData.email || !userData.password) {
            throw new Error("Tous les champs sont requis: nom d'utilisateur, email et mot de passe");
        }
        const response = await fetch(`${AUTH_URL}/users`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            mode: "cors",
            body: JSON.stringify(userData)
        });
        if (!response.ok) {
            const errorText = await response.text();
            console.error('User creation failed:', response.status, errorText);
            // Afficher en détail la réponse pour le débogage
            console.log('Response details:', {
                status: response.status,
                statusText: response.statusText,
                headers: Object.fromEntries(response.headers.entries()),
                text: errorText
            });
            let errorData;
            try {
                errorData = JSON.parse(errorText);
                console.log('Parsed error data:', errorData);
            } catch (e) {
                console.error('Failed to parse error response as JSON:', e);
                errorData = {
                    error: errorText || "Erreur inconnue"
                };
            }
            if (response.status === 409) {
                // Conflict - duplicate username or email
                throw new Error(errorData?.error || "Cet utilisateur existe déjà");
            } else if (response.status === 400) {
                // Bad request - validation error
                throw new Error(errorData?.error || "Données invalides. Vérifiez tous les champs requis.");
            }
            throw new Error(errorData?.error || "Failed to create user");
        }
        return response.json();
    } catch (error) {
        console.error("Error creating user:", error);
        throw error;
    }
}
}}),
"[project]/src/contexts/AuthContext.tsx [app-ssr] (ecmascript)": ((__turbopack_context__) => {
"use strict";

var { g: global, __dirname } = __turbopack_context__;
{
__turbopack_context__.s({
    "AuthProvider": (()=>AuthProvider),
    "useAuth": (()=>useAuth)
});
var __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2d$jsx$2d$dev$2d$runtime$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__ = __turbopack_context__.i("[project]/node_modules/next/dist/server/route-modules/app-page/vendored/ssr/react-jsx-dev-runtime.js [app-ssr] (ecmascript)");
var __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__ = __turbopack_context__.i("[project]/node_modules/next/dist/server/route-modules/app-page/vendored/ssr/react.js [app-ssr] (ecmascript)");
var __TURBOPACK__imported__module__$5b$project$5d2f$src$2f$utils$2f$api$2e$ts__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__ = __turbopack_context__.i("[project]/src/utils/api.ts [app-ssr] (ecmascript)");
"use client";
;
;
;
const AuthContext = /*#__PURE__*/ (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["createContext"])(undefined);
function AuthProvider({ children }) {
    const [user, setUser] = (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["useState"])(null);
    const [isLoading, setIsLoading] = (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["useState"])(true);
    const [error, setError] = (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["useState"])(null);
    // Check if user is already logged in (from localStorage)
    (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["useEffect"])(()=>{
        const storedUser = localStorage.getItem("user");
        if (storedUser) {
            try {
                setUser(JSON.parse(storedUser));
            } catch (e) {
                console.error("Failed to parse stored user", e);
                localStorage.removeItem("user");
            }
        }
        setIsLoading(false);
    }, []);
    // Login function
    const login = async (username, password)=>{
        setIsLoading(true);
        setError(null);
        try {
            const response = await (0, __TURBOPACK__imported__module__$5b$project$5d2f$src$2f$utils$2f$api$2e$ts__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["login"])(username, password);
            setUser(response.user);
            localStorage.setItem("user", JSON.stringify(response.user));
        } catch (err) {
            setError(err instanceof Error ? err.message : "Une erreur est survenue");
            throw err;
        } finally{
            setIsLoading(false);
        }
    };
    // Logout function
    const logout = ()=>{
        setUser(null);
        localStorage.removeItem("user");
    };
    const value = {
        user,
        isLoading,
        error,
        login,
        logout,
        isAuthenticated: !!user
    };
    return /*#__PURE__*/ (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2d$jsx$2d$dev$2d$runtime$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["jsxDEV"])(AuthContext.Provider, {
        value: value,
        children: children
    }, void 0, false, {
        fileName: "[project]/src/contexts/AuthContext.tsx",
        lineNumber: 68,
        columnNumber: 10
    }, this);
}
function useAuth() {
    const context = (0, __TURBOPACK__imported__module__$5b$project$5d2f$node_modules$2f$next$2f$dist$2f$server$2f$route$2d$modules$2f$app$2d$page$2f$vendored$2f$ssr$2f$react$2e$js__$5b$app$2d$ssr$5d$__$28$ecmascript$29$__["useContext"])(AuthContext);
    if (context === undefined) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
}
}}),
"[project]/node_modules/next/dist/server/route-modules/app-page/module.compiled.js [app-ssr] (ecmascript)": (function(__turbopack_context__) {

var { g: global, __dirname, m: module, e: exports } = __turbopack_context__;
{
"use strict";
if ("TURBOPACK compile-time falsy", 0) {
    "TURBOPACK unreachable";
} else {
    if ("TURBOPACK compile-time falsy", 0) {
        "TURBOPACK unreachable";
    } else {
        if ("TURBOPACK compile-time truthy", 1) {
            module.exports = __turbopack_context__.r("[externals]/next/dist/compiled/next-server/app-page.runtime.dev.js [external] (next/dist/compiled/next-server/app-page.runtime.dev.js, cjs)");
        } else {
            "TURBOPACK unreachable";
        }
    }
} //# sourceMappingURL=module.compiled.js.map
}}),
"[project]/node_modules/next/dist/server/route-modules/app-page/vendored/ssr/react-jsx-dev-runtime.js [app-ssr] (ecmascript)": (function(__turbopack_context__) {

var { g: global, __dirname, m: module, e: exports } = __turbopack_context__;
{
"use strict";
module.exports = __turbopack_context__.r("[project]/node_modules/next/dist/server/route-modules/app-page/module.compiled.js [app-ssr] (ecmascript)").vendored['react-ssr'].ReactJsxDevRuntime; //# sourceMappingURL=react-jsx-dev-runtime.js.map
}}),
"[project]/node_modules/next/dist/server/route-modules/app-page/vendored/ssr/react.js [app-ssr] (ecmascript)": (function(__turbopack_context__) {

var { g: global, __dirname, m: module, e: exports } = __turbopack_context__;
{
"use strict";
module.exports = __turbopack_context__.r("[project]/node_modules/next/dist/server/route-modules/app-page/module.compiled.js [app-ssr] (ecmascript)").vendored['react-ssr'].React; //# sourceMappingURL=react.js.map
}}),

};

//# sourceMappingURL=%5Broot%20of%20the%20server%5D__7e786d3b._.js.map