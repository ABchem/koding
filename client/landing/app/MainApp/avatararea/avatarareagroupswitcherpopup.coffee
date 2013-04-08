class AvatarPopupGroupSwitcher extends AvatarPopup

  viewAppended:->

    super

    @_popupList = new PopupList
      itemClass  : PopupGroupListItemWrapper

    @listController = new KDListViewController
      view                : @_popupList
      startWithLazyLoader : yes

    @listController.on "AvatarPopupShouldBeHidden", @bound 'hide'

    @avatarPopupContent.addSubView switchToTitle = new KDView
      height   : "auto"
      cssClass : "sublink top"
      partial  : "Switch to:"

    switchToTitle.addSubView new KDCustomHTMLView
      tagName    : 'span'
      cssClass   : 'icon help'
      tooltip    :
        title    : "Here you'll find the groups that you are a member of, clicking one of them will take you to a new browser tab."

    @avatarPopupContent.addSubView @listController.getView()

    @avatarPopupContent.addSubView new KDView
      height   : "auto"
      cssClass : "sublink"
      partial  : "<a href='#'>See all groups...</a>"
      click    : =>
        KD.getSingleton("appManager").open "Groups"
        @hide()

  accountChanged:->
    @listController.removeAllItems()

  populateGroups:->
    @listController.removeAllItems()
    @listController.showLazyLoader()
    KD.whoami().fetchGroupsIncludingPending (err, groups)=>
      if err then warn err
      else if groups?
        @listController.hideLazyLoader()
        @listController.addItem group  for group in groups

  show:->
    super
    @populateGroups()

class PopupGroupListItemWrapper extends KDListItemView

  constructor:(options = {}, data)->
    options.tagName or= "li"
    super

    {invitation} = @getData()
    if invitation
      @setClass 'pending'
      @child = new PopupGroupListItemPending {}, @getData()
    else
      @child = new PopupGroupListItem {}, @getData()

  viewAppended: JView::viewAppended

  pistachio: -> "{{> @child}}"

class PopupGroupListItem extends KDView

  constructor:(options = {}, data)->
    super

    {group:{title, avatar, slug}, roles} = @getData()

    roleClasses = roles.map((role)-> "role-#{role}").join ' '
    @setClass "role #{roleClasses}"

    @avatar = new KDCustomHTMLView
      tagName    : 'img'
      cssClass   : 'avatar-image'
      attributes :
        src      : avatar or "http://lorempixel.com/20/20?#{@utils.getRandomNumber()}"

    @switchLink = new CustomLinkView
      title       : title
      href        : "/#{if slug is 'koding' then '' else slug+'/'}Activity"
      target      : slug
      icon        :
        cssClass  : 'new-page'
        placement : 'right'
        tooltip   :
          title   : "Opens in a new browser window."
          delayIn : 300

  viewAppended: JView::viewAppended

  pistachio: ->
    {roles} = @getData()
    """
    <span class='avatar'>{{> @avatar}}</span>
    <div class='right-overflow'>
      {{> @switchLink}}
      <span class="roles">#{roles.join ', '}</span>
    </div>
    """

class PopupGroupListItemPending extends PopupGroupListItem

  constructor:(options = {}, data)->
    super

    {group:{title, slug}, invitation} = @getData()

    @acceptButton = new KDButtonView
      style       : 'clean-gray'
      title       : 'Accept Invitation'
      icon        : yes
      iconOnly    : yes
      iconClass   : 'accept'
      tooltip     :
        title     : 'Accept Invitation'
      callback    : =>
        invitation.acceptInvitationByInvitee (err)=>
          if err then warn err
          else
            @acceptButton.hide()
            @ignoreButton.hide()
            @parent.unsetClass 'pending'

    @ignoreButton = new KDButtonView
      style       : 'clean-gray'
      title       : 'Ignore Invitation'
      icon        : yes
      iconOnly    : yes
      iconClass   : 'ignore'
      tooltip     :
        title     : 'Ignore Invitation'
      callback    : =>
        invitation.ignoreInvitationByInvitee (err)=>
          if err then warn err
          else
            new KDNotificationView
              title    : 'Fair enough!'
              content  : 'If you change your mind, you can request access to the group anytime.'
              duration : 2000
            @parent.hide()

  viewAppended: JView::viewAppended

  pistachio: ->
    """
    <span class='avatar'>{{> @avatar}}</span>
    <div class='right-overflow'>
      <div class="buttons">
        {{> @acceptButton}}
        {{> @ignoreButton}}
      </div>
      {{> @switchLink}}
    </div>
    """
