// Copyright (c) 2017 The Hc developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

type newUserEmailTemplateData struct {
	Username string
	Link     string
	Email    string
}

type updateUserKeyEmailTemplateData struct {
	Link      string
	PublicKey string
	Email     string
}

type resetPasswordEmailTemplateData struct {
	Link  string
	Email string
}

type userLockedResetPasswordEmailTemplateData struct {
	Link  string
	Email string
}

type newProposalSubmittedTemplateData struct {
	Link     string
	Name     string
	Username string
	Email    string
}

type proposalEditedTemplateData struct {
	Link     string
	Name     string
	Version  string
	Username string
}

type proposalVoteStartedTemplateData struct {
	Link     string
	Name     string
	Username string
}

type proposalStatusChangeTemplateData struct {
	Link               string
	Name               string
	Username           string
	StatusChangeReason string
}

type proposalVoteAuthorizedTemplateData struct {
	Link     string
	Name     string
	Username string
	Email    string
}

type commentReplyOnProposalTemplateData struct {
	Commenter    string
	ProposalName string
	CommentLink  string
}

type commentReplyOnCommentTemplateData struct {
	Commenter    string
	ProposalName string
	CommentLink  string
}

const templateNewUserEmailRaw_old = `
Thanks for joining HcAutonomy, {{.Username}}!

Click the link below to verify your email and complete your registration:

{{.Link}}

You are receiving this email because {{.Email}} was used to register for HcAutonomy.
If you did not perform this action, please ignore this email.
`

const templateNewUserEmailRaw =`
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>

<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">For security purposes, please click the following
                    <a href="{{ .Link }}" style="color:#0e5bc9;">link</a> to verify your email address.</td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateResetPasswordEmailRaw_old = `
Click the link below to continue resetting your password:

{{.Link}}

You are receiving this email because a password reset was initiated for {{.Email}}
 on HcAutonomy. If you did not perform this action, please contact HcAutonomy
administrators.
`
const templateResetPasswordEmailRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">For security purposes, please click the following
                    <a href="{{ .Link }}" style="color:#0e5bc9;">link</a> to continue resetting your password.</td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateUpdateUserKeyEmailRaw_old = `
Click the link below to verify your new identity:

{{.Link}}

You are receiving this email because a new identity (public key: {{.PublicKey}})
was generated for {{.Email}} on HcAutonomy. If you did not perform this action,
please contact HcAutonomy administrators.
`

const templateUpdateUserKeyEmailRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">For security purposes, please click the following
                    <a href="{{ .Link }}" style="color:#0e5bc9;">link</a> to verify your new identity.</td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateUserLockedResetPasswordRaw_old = `
Your account was locked due to too many login attempts. You need to reset your
password in order to unlock your account:

{{.Link}}

You are receiving this email because someone made too many login attempts for
{{.Email}} on HcAutonomy. If that was not you, please notify HcAutonomy administrators.
`
const templateUserLockedResetPasswordRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Your account was locked due to too many login attempts. You need to reset your password in order to unlock your account. For security purposes, please click the following
                    <a href="{{ .Link }}" style="color:#0e5bc9;">link</a> to continue resetting your password.</td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateNewProposalSubmittedRaw_old = `
A new proposal has been submitted on HcAutonomy by {{.Username}} ({{.Email}}):

{{.Name}}
{{.Link}}
`
const templateNewProposalSubmittedRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">A new proposal has been submitted on HcAutonomy by {{.Username}} ({{.Email}})
                    <a href="{{ .Link }}" style="color:#0e5bc9;">link</a> </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalVettedRaw_old = `
A new proposal has just been approved on HcAutonomy, authored by {{.Username}}:

{{.Name}}
{{.Link}}
`
const templateProposalVettedRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">A new proposal has just been approved on HcAutonomy, authored by {{.Username}}
                    </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalEditedRaw_old = `
A proposal by {{.Username}} has just been edited:

{{.Name}} (Version: {{.Version}})
{{.Link}}
`
const templateProposalEditedRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">A proposal by {{.Username}} has just been edited:
                     </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalVoteStartedRaw_old = `
Voting has started for the following proposal on HcAutonomy, authored by {{.Username}}:

{{.Name}}
{{.Link}}
`
const templateProposalVoteStartedRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Voting has started for the following proposal on HcAutonomy, authored by {{.Username}}
                     </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalVoteAuthorizedRaw_old = `
Voting has been authorized for the following proposal on HcAutonomy by {{.Username}} ({{.Email}}):

{{.Name}}
{{.Link}}
`
const templateProposalVoteAuthorizedRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Voting has been authorized for the following proposal on HcAutonomy by {{.Username}} ({{.Email}}):
                    </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalVettedForAuthorRaw_old = `
Your proposal has just been approved on HcAutonomy!

You will need to authorize a proposal vote before an administrator will be
allowed to start the voting period on your proposal.  You can authorize a
proposal vote by opening the proposal page and clicking on the "Authorize
Voting to Start" button.

You must authorize a proposal vote within 14 days.  If you fail to do so, your
proposal will be considered abandoned.

{{.Name}}
{{.Link}}
`
const templateProposalVettedForAuthorRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Your proposal has just been approved on HcAutonomy!

You will need to authorize a proposal vote before an administrator will be
allowed to start the voting period on your proposal.  You can authorize a
proposal vote by opening the proposal page and clicking on the "Authorize
Voting to Start" button.

You must authorize a proposal vote within 14 days.  If you fail to do so, your
proposal will be considered abandoned.
                     </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalCensoredForAuthorRaw_old = `
Your proposal on HcAutonomy has been censored:

{{.Name}}
{{.Link}}
Reason: {{.StatusChangeReason}}
`
const templateProposalCensoredForAuthorRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Your proposal on HcAutonomy has been censored:
                    </td>
                </tr>
				<tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Reason: {{.StatusChangeReason}} </td>
                </tr>
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateProposalVoteStartedForAuthorRaw_old = `
Voting has just started for your proposal on HcAutonomy!

{{.Name}}
{{.Link}}
`
const templateProposalVoteStartedForAuthorRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Voting has just started for your proposal on HcAutonomy!
                    <a href="{{ .Link }}" style="color:#0e5bc9;">link</a> </td>
                </tr>

                    <tr>
                  <td style="font-size:16px;color:#1B222D;">
                    <a href="{{ .Link }}" style="color:#0e5bc9;">{{ .Link }}</a>
                  </td>
                </tr>

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateCommentReplyOnProposalRaw_old = `
{{.Commenter}} has commented on your proposal!

Proposal: {{.ProposalName}}
Comment: {{.CommentLink}}
`
const templateCommentReplyOnProposalRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">{{.Commenter}} has commented on your proposal!
                    </td>
                </tr>
				
				<tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Proposal: {{.ProposalName}}
                     </td>
                </tr>
				
				<tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Proposal: {{.ProposalName}}
                     </td>
                </tr>
				
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">Comment: 
                    <a href="{{ .CommentLink }}" style="color:#0e5bc9;"> CommentLink</a>
{{.CommentLink}}
                  </td>
                </tr>
				

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`

const templateCommentReplyOnCommentRaw_old = `
{{.Commenter}} has replied to your comment!

Proposal: {{.ProposalName}}
Comment: {{.CommentLink}}
`
const templateCommentReplyOnCommentRaw = `
<!doctype html>
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title>EDM</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>


<body>
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tbody>
        <tr>
          <td style="font-family: Arial, sans-serif;">
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="600px" style="text-align:center;">
              <tbody>
                <tr>
                  <td>
                    <div style="width:100%;max-width: 100%;text-align:center"><img style="width: 600px;max-width: 100%; display: block;margin:auto" src="http://test.legenddigital.com.au/jack/hcash_email/images/banner.jpg" width="600px" alt=""></div>
                  </td>
                </tr>
                <tr>
                  <td style="font-size:24px;color:#1B222D;padding:30px 10px 30px 10px;">
                    Dear HCASH user,
                  </td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">{{.Commenter}} has replied to your comment!
                    </td>
                </tr>
				
				<tr>
                  <td style="font-size:18px;color:#1B222D;padding:0 10px 80px 10px;">Proposal: {{.ProposalName}}
                     </td>
                </tr>
				
				
                    <tr>
                  <td style="font-size:16px;color:#1B222D;">Comment:
                    <a href="{{ .CommentLink }}" style="color:#0e5bc9;"> CommentLink</a>
					{{.CommentLink}}
                  </td>
                </tr>
				

                <tr>
                  <td style="border-top: 1px solid #BABCC0; font-size:18px;color:#1B222D;padding:50px 10px 20px 10px;">Yours Sincerely</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 30px 10px;">HCASH Team</td>
                </tr>
                <tr>
                  <td style="font-size:18px;color:#1B222D;padding:0px 10px 120px 10px;"><img src="http://test.legenddigital.com.au/jack/hcash_email/images/icon.jpg" width="46"></td>
                </tr>
                <tr>
                  <td style="background-color: #f7f6f7;color: #000000;font-size:14px; text-align:center;padding:20px 10px;">© Copyright 2018
                    <span style="color:#8219bf;">HCASH</span>, All rights reserved.</td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </body>
</html>
`