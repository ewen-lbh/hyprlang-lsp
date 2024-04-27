package parser

import "image/color"

type Configuration struct {
	CustomVariables map[string]string

	General ConfigurationGeneral
	Decoration ConfigurationDecoration
	Animations ConfigurationAnimations
	Input ConfigurationInput
	Gestures ConfigurationGestures
	Group ConfigurationGroup
	Misc ConfigurationMisc
	Binds ConfigurationBinds
	XWayland ConfigurationXWayland
	OpenGL ConfigurationOpenGL
	Debug ConfigurationDebug
}


type ConfigurationGeneral struct {
	// mouse sensitivity (legacy, may cause bugs if not 1, prefer input:sensitivity)
	Sensitivity float32 `json:"sensitivity"`

	// size of the border around windows
	BorderSize int `json:"border_size"`

	// disable borders for floating windows
	NoBorderOnFloating bool `json:"no_border_on_floating"`

	// gaps between windows, also supports css style gaps (top, right, bottom, left -> 5,10,15,20)
	GapsIn int `json:"gaps_in"`

	// gaps between windows and monitor edges, also supports css style gaps (top, right, bottom, left -> 5,10,15,20)
	GapsOut int `json:"gaps_out"`

	// gaps between workspaces. Stacks with gaps_out.
	GapsWorkspaces int `json:"gaps_workspaces"`

	// border color for inactive windows
	ColInactiveBorder GradientValue `json:"col.inactive_border"`

	// border color for the active window
	ColActiveBorder GradientValue `json:"col.active_border"`

	// inactive border color for window that cannot be added to a group (see denywindowfromgroup dispatcher)
	ColNogroupBorder GradientValue `json:"col.nogroup_border"`

	// active border color for window that cannot be added to a group
	ColNogroupBorderActive GradientValue `json:"col.nogroup_border_active"`

	// in seconds, after how many seconds of cursor's inactivity to hide it. Set to 0 for never.
	CursorInactiveTimeout int `json:"cursor_inactive_timeout"`

	// which layout to use. [dwindle/master]
	Layout string `json:"layout"`

	// if true, will not warp the cursor in many cases (focusing, keybinds, etc)
	NoCursorWarps bool `json:"no_cursor_warps"`

	// if true, will not fall back to the next available window when moving focus in a direction where no window was found
	NoFocusFallback bool `json:"no_focus_fallback"`

	// if on, will also apply the sensitivity to raw mouse output (e.g. sensitivity in games) NOTICE: really not recommended.
	ApplySensToRaw bool `json:"apply_sens_to_raw"`

	// enables resizing windows by clicking and dragging on borders and gaps
	ResizeOnBorder bool `json:"resize_on_border"`

	// extends the area around the border where you can click and drag on, only used when general:resize_on_border is on.
	ExtendBorderGrabArea int `json:"extend_border_grab_area"`

	// show a cursor icon when hovering over borders, only used when general:resize_on_border is on.
	HoverIconOnBorder bool `json:"hover_icon_on_border"`

	// master switch for allowing tearing to occur. See the Tearing page.
	AllowTearing bool `json:"allow_tearing"`

	// force floating windows to use a specific corner when being resized (1-4 going clockwise from top left, 0 to disable)
	ResizeCorner int `json:"resize_corner"`

}


type ConfigurationDecoration struct {
	// rounded corners' radius (in layout px)
	Rounding int `json:"rounding"`

	// opacity of active windows. [0.0 - 1.0]
	ActiveOpacity float32 `json:"active_opacity"`

	// opacity of inactive windows. [0.0 - 1.0]
	InactiveOpacity float32 `json:"inactive_opacity"`

	// opacity of fullscreen windows. [0.0 - 1.0]
	FullscreenOpacity float32 `json:"fullscreen_opacity"`

	// enable drop shadows on windows
	DropShadow bool `json:"drop_shadow"`

	// Shadow range ("size") in layout px
	ShadowRange int `json:"shadow_range"`

	// in what power to render the falloff (more power, the faster the falloff) [1 - 4]
	ShadowRenderPower int `json:"shadow_render_power"`

	// if true, the shadow will not be rendered behind the window itself, only around it.
	ShadowIgnoreWindow bool `json:"shadow_ignore_window"`

	// shadow's color. Alpha dictates shadow's opacity.
	ColShadow color.RGBA `json:"col.shadow"`

	// inactive shadow color. (if not set, will fall back to col.shadow)
	ColShadowInactive color.RGBA `json:"col.shadow_inactive"`

	// shadow's rendering offset.
	ShadowOffset [2]float32 `json:"shadow_offset"`

	// shadow's scale. [0.0 - 1.0]
	ShadowScale float32 `json:"shadow_scale"`

	// enables dimming of inactive windows
	DimInactive bool `json:"dim_inactive"`

	// how much inactive windows should be dimmed [0.0 - 1.0]
	DimStrength float32 `json:"dim_strength"`

	// how much to dim the rest of the screen by when a special workspace is open. [0.0 - 1.0]
	DimSpecial float32 `json:"dim_special"`

	// how much the dimaround window rule should dim by. [0.0 - 1.0]
	DimAround float32 `json:"dim_around"`

	// a path to a custom shader to be applied at the end of rendering. See examples/screenShader.frag for an example.
	ScreenShader string `json:"screen_shader"`

	Blur ConfigurationDecorationBlur `json:"blur"`
}

type ConfigurationDecorationBlur struct {
	// enable kawase window background blur
	Enabled bool `json:"enabled"`

	// blur size (distance)
	Size int `json:"size"`

	// the amount of passes to perform
	Passes int `json:"passes"`

	// make the blur layer ignore the opacity of the window
	IgnoreOpacity bool `json:"ignore_opacity"`

	// whether to enable further optimizations to the blur. Recommended to leave on, as it will massively improve performance.
	NewOptimizations bool `json:"new_optimizations"`

	// if enabled, floating windows will ignore tiled windows in their blur. Only available if blur_new_optimizations is true. Will reduce overhead on floating blur significantly.
	Xray bool `json:"xray"`

	// how much noise to apply. [0.0 - 1.0]
	Noise float32 `json:"noise"`

	// contrast modulation for blur. [0.0 - 2.0]
	Contrast float32 `json:"contrast"`

	// brightness modulation for blur. [0.0 - 2.0]
	Brightness float32 `json:"brightness"`

	// Increase saturation of blurred colors. [0.0 - 1.0]
	Vibrancy float32 `json:"vibrancy"`

	// How strong the effect of vibrancy is on dark areas . [0.0 - 1.0]
	VibrancyDarkness float32 `json:"vibrancy_darkness"`

	// whether to blur behind the special workspace (note: expensive)
	Special bool `json:"special"`

	// whether to blur popups (e.g. right-click menus)
	Popups bool `json:"popups"`

	// works like ignorealpha in layer rules. If pixel opacity is below set value, will not blur. [0.0 - 1.0]
	PopupsIgnorealpha float32 `json:"popups_ignorealpha"`

}


type ConfigurationAnimations struct {
	// enable animations
	Enabled bool `json:"enabled"`

	// enable first launch animation
	FirstLaunchAnimation bool `json:"first_launch_animation"`

}


type ConfigurationInput struct {
	// Appropriate XKB keymap parameter. See the note below.
	KbModel string `json:"kb_model"`

	// Appropriate XKB keymap parameter
	KbLayout string `json:"kb_layout"`

	// Appropriate XKB keymap parameter
	KbVariant string `json:"kb_variant"`

	// Appropriate XKB keymap parameter
	KbOptions string `json:"kb_options"`

	// Appropriate XKB keymap parameter
	KbRules string `json:"kb_rules"`

	// If you prefer, you can use a path to your custom .xkb file.
	KbFile string `json:"kb_file"`

	// Engage numlock by default.
	NumlockByDefault bool `json:"numlock_by_default"`

	// Determines how keybinds act when multiple layouts are used. If false, keybinds will always act as if the first specified layout is active. If true, keybinds specified by symbols are activated when you type the respective symbol with the current layout.
	ResolveBindsBySym bool `json:"resolve_binds_by_sym"`

	// The repeat rate for held-down keys, in repeats per second.
	RepeatRate int `json:"repeat_rate"`

	// Delay before a held-down key is repeated, in milliseconds.
	RepeatDelay int `json:"repeat_delay"`

	// Sets the mouse input sensitivity. Value is clamped to the range -1.0 to 1.0. libinput#pointer-acceleration
	Sensitivity float32 `json:"sensitivity"`

	// Sets the cursor acceleration profile. Can be one of adaptive, flat. Can also be custom, see below. Leave empty to use libinput's default mode for your input device. libinput#pointer-acceleration [adaptive/flat/custom]
	AccelProfile string `json:"accel_profile"`

	// Force no cursor acceleration. This bypasses most of your pointer settings to get as raw of a signal as possible. Enabling this is not recommended due to potential cursor desynchronization.
	ForceNoAccel bool `json:"force_no_accel"`

	// Switches RMB and LMB
	LeftHanded bool `json:"left_handed"`

	// Sets the scroll acceleration profile, when accel_profile is set to custom. Has to be in the form <step> <points>. Leave empty to have a flat scroll curve.
	ScrollPoints string `json:"scroll_points"`

	// Sets the scroll method. Can be one of 2fg (2 fingers), edge, on_button_down, no_scroll. libinput#scrolling [2fg/edge/on_button_down/no_scroll]
	ScrollMethod string `json:"scroll_method"`

	// Sets the scroll button. Has to be an int, cannot be a string. Check wev if you have any doubts regarding the ID. 0 means default.
	ScrollButton int `json:"scroll_button"`

	// If the scroll button lock is enabled, the button does not need to be held down. Pressing and releasing the button toggles the button lock, which logically holds the button down or releases it. While the button is logically held down, motion events are converted to scroll events.
	ScrollButtonLock bool `json:"scroll_button_lock"`

	// Multiplier added to scroll movement for external mice. Note that there is a separate setting for touchpad scroll_factor.
	ScrollFactor float32 `json:"scroll_factor"`

	// Inverts scrolling direction. When enabled, scrolling moves content directly, rather than manipulating a scrollbar.
	NaturalScroll bool `json:"natural_scroll"`

	// Specify if and how cursor movement should affect window focus. See the note below. [0/1/2/3]
	FollowMouse int `json:"follow_mouse"`

	// If disabled, mouse focus won't switch to the hovered window unless the mouse crosses a window boundary when follow_mouse=1.
	MouseRefocus bool `json:"mouse_refocus"`

	// If enabled (1 or 2), focus will change to the window under the cursor when changing from tiled-to-floating and vice versa. If 2, focus will also follow mouse on float-to-float switches.
	FloatSwitchOverrideFocus int `json:"float_switch_override_focus"`

	// if enabled, having only floating windows in the special workspace will not block focusing windows in the regular workspace.
	SpecialFallthrough bool `json:"special_fallthrough"`

	// Handles axis events around (gaps/border for tiled, dragarea/border for floated) a focused window. 0 ignores axis events 1 sends out-of-bound coordinates 2 fakes pointer coordinates to the closest point inside the window 3 warps the cursor to the closest point inside the window
	OffWindowAxisEvents int `json:"off_window_axis_events"`

}


type ConfigurationGestures struct {
	// enable workspace swipe gesture on touchpad
	WorkspaceSwipe bool `json:"workspace_swipe"`

	// how many fingers for the touchpad gesture
	WorkspaceSwipeFingers int `json:"workspace_swipe_fingers"`

	// in px, the distance of the touchpad gesture
	WorkspaceSwipeDistance int `json:"workspace_swipe_distance"`

	// enable workspace swiping from the edge of a touchscreen
	WorkspaceSwipeTouch bool `json:"workspace_swipe_touch"`

	// invert the direction
	WorkspaceSwipeInvert bool `json:"workspace_swipe_invert"`

	// minimum speed in px per timepoint to force the change ignoring cancel_ratio. Setting to 0 will disable this mechanic.
	WorkspaceSwipeMinSpeedToForce int `json:"workspace_swipe_min_speed_to_force"`

	// how much the swipe has to proceed in order to commence it. (0.7 -> if > 0.7 * distance, switch, if less, revert) [0.0 - 1.0]
	WorkspaceSwipeCancelRatio float32 `json:"workspace_swipe_cancel_ratio"`

	// whether a swipe right on the last workspace should create a new one.
	WorkspaceSwipeCreateNew bool `json:"workspace_swipe_create_new"`

	// if enabled, switching direction will be locked when you swipe past the direction_lock_threshold (touchpad only).
	WorkspaceSwipeDirectionLock bool `json:"workspace_swipe_direction_lock"`

	// in px, the distance to swipe before direction lock activates (touchpad only).
	WorkspaceSwipeDirectionLockThreshold int `json:"workspace_swipe_direction_lock_threshold"`

	// if enabled, swiping will not clamp at the neighboring workspaces but continue to the further ones.
	WorkspaceSwipeForever bool `json:"workspace_swipe_forever"`

	// if enabled, swiping will use the r prefix instead of the m prefix for finding workspaces.
	WorkspaceSwipeUseR bool `json:"workspace_swipe_use_r"`

}


type ConfigurationGroup struct {
	// whether new windows in a group spawn after current or at group tail
	InsertAfterCurrent bool `json:"insert_after_current"`

	// whether Hyprland should focus on the window that has just been moved out of the group
	FocusRemovedWindow bool `json:"focus_removed_window"`

	// active group border color
	ColBorderActive GradientValue `json:"col.border_active"`

	// inactive (out of focus) group border color
	ColBorderInactive GradientValue `json:"col.border_inactive"`

	// active locked group border color
	ColBorderLockedActive GradientValue `json:"col.border_locked_active"`

	// inactive locked group border color
	ColBorderLockedInactive GradientValue `json:"col.border_locked_inactive"`

	Groupbar ConfigurationGroupGroupbar `json:"groupbar"`
}

type ConfigurationGroupGroupbar struct {
	// enables groupbars
	Enabled bool `json:"enabled"`

	// font used to display groupbar titles
	FontFamily string `json:"font_family"`

	// font size of groupbar title
	FontSize int `json:"font_size"`

	// enables gradients
	Gradients bool `json:"gradients"`

	// height of the groupbar
	Height int `json:"height"`

	// sets the decoration priority for groupbars
	Priority int `json:"priority"`

	// whether to render titles in the group bar decoration
	RenderTitles bool `json:"render_titles"`

	// whether scrolling in the groupbar changes group active window
	Scrolling bool `json:"scrolling"`

	// controls the group bar text color
	TextColor color.RGBA `json:"text_color"`

	// active group border color
	ColActive GradientValue `json:"col.active"`

	// inactive (out of focus) group border color
	ColInactive GradientValue `json:"col.inactive"`

	// active locked group border color
	ColLockedActive GradientValue `json:"col.locked_active"`

	// inactive locked group border color
	ColLockedInactive GradientValue `json:"col.locked_inactive"`

}


type ConfigurationMisc struct {
	// disables the random Hyprland logo / anime girl background. :(
	DisableHyprlandLogo bool `json:"disable_hyprland_logo"`

	// disables the Hyprland splash rendering. (requires a monitor reload to take effect)
	DisableSplashRendering bool `json:"disable_splash_rendering"`

	// Changes the color of the splash text (requires a monitor reload to take effect).
	ColSplash color.RGBA `json:"col.splash"`

	// Changes the font used to render the splash text, selected from system fonts (requires a monitor reload to take effect).
	SplashFontFamily string `json:"splash_font_family"`

	// Enforce any of the 3 default wallpapers. Setting this to 0 or 1 disables the anime background. -1 means "random". [-1/0/1/2]
	ForceDefaultWallpaper int `json:"force_default_wallpaper"`

	// controls the VFR status of Hyprland. Heavily recommended to leave enabled to conserve resources.
	Vfr bool `json:"vfr"`

	// controls the VRR (Adaptive Sync) of your monitors. 0 - off, 1 - on, 2 - fullscreen only [0/1/2]
	Vrr int `json:"vrr"`

	// If DPMS is set to off, wake up the monitors if the mouse moves.
	MouseMoveEnablesDpms bool `json:"mouse_move_enables_dpms"`

	// If DPMS is set to off, wake up the monitors if a key is pressed.
	KeyPressEnablesDpms bool `json:"key_press_enables_dpms"`

	// Will make mouse focus follow the mouse when drag and dropping. Recommended to leave it enabled, especially for people using focus follows mouse at 0.
	AlwaysFollowOnDnd bool `json:"always_follow_on_dnd"`

	// If true, will make keyboard-interactive layers keep their focus on mouse move (e.g. wofi, bemenu)
	LayersHogKeyboardFocus bool `json:"layers_hog_keyboard_focus"`

	// If true, will animate manual window resizes/moves
	AnimateManualResizes bool `json:"animate_manual_resizes"`

	// If true, will animate windows being dragged by mouse, note that this can cause weird behavior on some curves
	AnimateMouseWindowdragging bool `json:"animate_mouse_windowdragging"`

	// If true, the config will not reload automatically on save, and instead needs to be reloaded with hyprctl reload. Might save on battery.
	DisableAutoreload bool `json:"disable_autoreload"`

	// Enable window swallowing
	EnableSwallow bool `json:"enable_swallow"`

	// The class regex to be used for windows that should be swallowed (usually, a terminal). To know more about the list of regex which can be used use this cheatsheet.
	SwallowRegex string `json:"swallow_regex"`

	// The title regex to be used for windows that should not be swallowed by the windows specified in swallow_regex  (e.g. wev). The regex is matched against the parent (e.g. Kitty) window's title on the assumption that it changes to whatever process it's running.
	SwallowExceptionRegex string `json:"swallow_exception_regex"`

	// Whether Hyprland should focus an app that requests to be focused (an activate request)
	FocusOnActivate bool `json:"focus_on_activate"`

	// Disables direct scanout. Direct scanout attempts to reduce lag when there is only one fullscreen application on a screen (e.g. game). It is also recommended to set this to true if the fullscreen application shows graphical glitches.
	NoDirectScanout bool `json:"no_direct_scanout"`

	// Hides the cursor when the last input was a touch input until a mouse input is done.
	HideCursorOnTouch bool `json:"hide_cursor_on_touch"`

	// Hides the cursor when you press any key until the mouse is moved.
	HideCursorOnKeyPress bool `json:"hide_cursor_on_key_press"`

	// Whether mouse moving into a different monitor should focus it
	MouseMoveFocusesMonitor bool `json:"mouse_move_focuses_monitor"`

	// disables warnings about incompatible portal implementations.
	SuppressPortalWarnings bool `json:"suppress_portal_warnings"`

	// [Warning: buggy] starts rendering before your monitor displays a frame in order to lower latency
	RenderAheadOfTime bool `json:"render_ahead_of_time"`

	// how many ms of safezone to add to rendering ahead of time. Recommended 1-2.
	RenderAheadSafezone int `json:"render_ahead_safezone"`

	// the factor to zoom by around the cursor. Like a magnifying glass. Minimum 1.0 (meaning no zoom)
	CursorZoomFactor float32 `json:"cursor_zoom_factor"`

	// whether the zoom should follow the cursor rigidly (cursor is always centered if it can be) or loosely
	CursorZoomRigid bool `json:"cursor_zoom_rigid"`

	// if true, will allow you to restart a lockscreen app in case it crashes (red screen of death)
	AllowSessionLockRestore bool `json:"allow_session_lock_restore"`

	// change the background color. (requires enabled disable_hyprland_logo)
	BackgroundColor color.RGBA `json:"background_color"`

	// close the special workspace if the last window is removed
	CloseSpecialOnEmpty bool `json:"close_special_on_empty"`

	// if there is a fullscreen window, whether a new tiled window opened should replace the fullscreen one or stay behind. 0 - behind, 1 - takes over, 2 - unfullscreen the current fullscreen window [0/1/2]
	NewWindowTakesOverFullscreen int `json:"new_window_takes_over_fullscreen"`

	// whether to enable hyprcursor support
	EnableHyprcursor bool `json:"enable_hyprcursor"`

	// if enabled, windows will open on the workspace they were invoked on. 0 - disabled, 1 - single-shot, 2 - persistent (all children too)
	InitialWorkspaceTracking int `json:"initial_workspace_tracking"`

}


type ConfigurationBinds struct {
	// if disabled, will not pass the mouse events to apps / dragging windows around if a keybind has been triggered.
	PassMouseWhenBound bool `json:"pass_mouse_when_bound"`

	// in ms, how many ms to wait after a scroll event to allow passing another one for the binds.
	ScrollEventDelay int `json:"scroll_event_delay"`

	// If enabled, an attempt to switch to the currently focused workspace will instead switch to the previous workspace. Akin to i3's auto_back_and_forth.
	WorkspaceBackAndForth bool `json:"workspace_back_and_forth"`

	// If enabled, workspaces don't forget their previous workspace, so cycles can be created by switching to the first workspace in a sequence, then endlessly going to the previous workspace.
	AllowWorkspaceCycles bool `json:"allow_workspace_cycles"`

	// Whether switching workspaces should center the cursor on the workspace (0) or on the last active window for that workspace (1)
	WorkspaceCenterOn int `json:"workspace_center_on"`

	// sets the preferred focus finding method when using focuswindow/movewindow/etc with a direction. 0 - history (recent have priority), 1 - length (longer shared edges have priority)
	FocusPreferredMethod int `json:"focus_preferred_method"`

	// If enabled, dispatchers like moveintogroup, moveoutofgroup and movewindoworgroup will ignore lock per group.
	IgnoreGroupLock bool `json:"ignore_group_lock"`

	// If enabled, when on a fullscreen window, movefocus will cycle fullscreen, if not, it will move the focus in a direction.
	MovefocusCyclesFullscreen bool `json:"movefocus_cycles_fullscreen"`

	// If enabled, apps that request keybinds to be disabled (e.g. VMs) will not be able to do so.
	DisableKeybindGrabbing bool `json:"disable_keybind_grabbing"`

}


type ConfigurationXWayland struct {
	// uses the nearest neigbor filtering for xwayland apps, making them pixelated rather than blurry
	UseNearestNeighbor bool `json:"use_nearest_neighbor"`

	// forces a scale of 1 on xwayland windows on scaled displays.
	ForceZeroScaling bool `json:"force_zero_scaling"`

}


type ConfigurationOpenGL struct {
	// reduces flickering on nvidia at the cost of possible frame drops on lower-end GPUs. On non-nvidia, this is ignored.
	NvidiaAntiFlicker bool `json:"nvidia_anti_flicker"`

	// forces introspection at all times. Introspection is aimed at reducing GPU usage in certain cases, but might cause graphical glitches on nvidia. 0 - nothing, 1 - force always on, 2 - force always on if nvidia
	ForceIntrospection int `json:"force_introspection"`

}


type ConfigurationDebug struct {
	// print the debug performance overlay. Disable VFR for accurate results.
	Overlay bool `json:"overlay"`

	// (epilepsy warning!) flash areas updated with damage tracking
	DamageBlink bool `json:"damage_blink"`

	// disable logging to a file
	DisableLogs bool `json:"disable_logs"`

	// disables time logging
	DisableTime bool `json:"disable_time"`

	// redraw only the needed bits of the display. Do not change. (default: full - 2) monitor - 1, none - 0
	DamageTracking int `json:"damage_tracking"`

	// enables logging to stdout
	EnableStdoutLogs bool `json:"enable_stdout_logs"`

	// set to 1 and then back to 0 to crash Hyprland.
	ManualCrash int `json:"manual_crash"`

	// if true, do not display config file parsing errors.
	SuppressErrors bool `json:"suppress_errors"`

	// sets the timeout in seconds for watchdog to abort processing of a signal of the main thread. Set to 0 to disable.
	WatchdogTimeout int `json:"watchdog_timeout"`

	// disables verification of the scale factors. Will result in pixel alignment and rounding errors.
	DisableScaleChecks bool `json:"disable_scale_checks"`

	// limits the number of displayed config file parsing errors.
	ErrorLimit int `json:"error_limit"`

}




