package api

func EmailVerification(email, code string) string {
	return `
	<div>
	  <table
		width="100%"
		border="0"
		cellspacing="0"
		cellpadding="0"
		style="width: 100% !important"
	  >
		<tbody>
		  <tr>
			<td align="center">
			  <table
				style="
				  border: 1px solid #eaeaea;
				  border-radius: 5px;
				  margin: 40px 0;
				"
				width="600"
				border="0"
				cellspacing="0"
				cellpadding="40"
			  >
				<tbody>
				  <tr>
					<td align="center">
					  <div style="text-align: left; width: 465px">
						<table
						  width="100%"
						  border="0"
						  cellspacing="0"
						  cellpadding="0"
						  style="width: 100% !important"
						>
						  <tbody>
							<tr>
							  <td align="center">
								<img
								  src="$LOGO_URL"
								  alt="$PASSWORD_MANAGER_NAME"
								/>
								<h1
								  style="
									color: #000;
									font-family: -apple-system, BlinkMacSystemFont,
									  'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
									  'Cantarell', 'Fira Sans', 'Droid Sans',
									  'Helvetica Neue', sans-serif;
									font-size: 24px;
									font-weight: normal;
									margin: 30px 0;
									padding: 0;
								  "
								>
								  Here's your $PASSWORD_MANAGER_NAME Verification Code
								</h1>
							  </td>
							</tr>
						  </tbody>
						</table>

						<p
						  style="
							color: #000;
							font-family: -apple-system, BlinkMacSystemFont,
							  'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
							  'Cantarell', 'Fira Sans', 'Droid Sans',
							  'Helvetica Neue', sans-serif;
							font-size: 14px;
							line-height: 24px;
						  "
						>
						  Hi <strong>` + email + `</strong>,
						</p>
						<p
						  style="
							color: #000;
							font-family: -apple-system, BlinkMacSystemFont,
							  'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
							  'Cantarell', 'Fira Sans', 'Droid Sans',
							  'Helvetica Neue', sans-serif;
							font-size: 14px;
							line-height: 24px;
						  "
						>
						  Continue signing up for $PASSWORD_MANAGER_NAME by entering the code
						  below:
						</p>
						<br />

						<table
						  width="100%"
						  border="0"
						  cellspacing="0"
						  cellpadding="0"
						  style="width: 100% !important"
						>
						  <tbody>
							<tr>
							  <td
								align="center"
								valign="middle"
								style="
								  color: #eeeeee !important;
								  background-color: #1163e6;
								  border-radius: 5px;
								  padding: 5px;
								  font: 300 20px 'SFMono-Regular', Consolas,
									'Liberation Mono', Menlo, monospace;
								"
							  >
								` + code + `
							  </td>
							</tr>
						  </tbody>
						</table>

						<br />

						<hr
						  style="
							border: none;
							border-top: 1px solid #eaeaea;
							margin: 26px 0;
							width: 100%;
						  "
						/>
						<p
						  style="
							color: #666666;
							font-family: -apple-system, BlinkMacSystemFont,
							  'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
							  'Cantarell', 'Fira Sans', 'Droid Sans',
							  'Helvetica Neue', sans-serif;
							font-size: 12px;
							line-height: 24px;
						  "
						>
						  &copy; $PASSWORD_MANAGER_NAME. All rights reserved.
						</p>
					  </div>
					</td>
				  </tr>
				</tbody>
			  </table>
			</td>
		  </tr>
		</tbody>
	  </table>
	</div>
	`
}

func ResetMasterPasswordInstructions(email, code string) string {
	return `
	<div>
		<div>
			<table
				width="100%"
				border="0"
				cellspacing="0"
				cellpadding="0"
				style="width: 100% !important"
			>
				<tbody>
					<tr>
					<td align="center">
						<table
						style="
							border: 1px solid #eaeaea;
							border-radius: 5px;
							margin: 40px 0;
						"
						width="600"
						border="0"
						cellspacing="0"
						cellpadding="40"
						>
						<tbody>
							<tr>
							<td align="center">
								<div style="text-align: left; width: 465px">
								<table
									width="100%"
									border="0"
									cellspacing="0"
									cellpadding="0"
									style="width: 100% !important"
								>
									<tbody>
									<tr>
										<td align="center">
										<img
											src="$LOGO_URL"
											alt="$PASSWORD_MANAGER_NAME"
										/>
										<h1
											style="
											color: #000;
											font-family: -apple-system, BlinkMacSystemFont,
												'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
												'Cantarell', 'Fira Sans', 'Droid Sans',
												'Helvetica Neue', sans-serif;
											font-size: 24px;
											font-weight: normal;
											margin: 30px 0;
											padding: 0;
											"
										>
											Reset master password request for ` + email + `
										</h1>
										</td>
									</tr>
									</tbody>
								</table>

								<p
									style="
									color: #000;
									font-family: -apple-system, BlinkMacSystemFont,
										'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
										'Cantarell', 'Fira Sans', 'Droid Sans',
										'Helvetica Neue', sans-serif;
									font-size: 14px;
									line-height: 24px;
									"
								>
									A request was made to change the master password for
									this account. If you didn’t intend to change your master
									password you can ignore this email to leave it
									unchanged.

									<br />
									<br />

									this is your reseting master password code:
								</p>
								<br />

								<table
									width="100%"
									border="0"
									cellspacing="0"
									cellpadding="0"
									style="width: 100% !important"
								>
									<tbody>
									<tr>
										<td
										align="center"
										valign="middle"
										style="
											color: #eeeeee !important;
											background-color: #1163e6;
											border-radius: 5px;
											padding: 5px;
											font: 300 20px 'SFMono-Regular', Consolas,
											'Liberation Mono', Menlo, monospace;
										"
										>
										` + code + `
										</td>
									</tr>
									</tbody>
								</table>

								<br />

								<hr
									style="
									border: none;
									border-top: 1px solid #eaeaea;
									margin: 26px 0;
									width: 100%;
									"
								/>
								<p
									style="
									color: #666666;
									font-family: -apple-system, BlinkMacSystemFont,
										'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
										'Cantarell', 'Fira Sans', 'Droid Sans',
										'Helvetica Neue', sans-serif;
									font-size: 12px;
									line-height: 24px;
									"
								>
									&copy; $PASSWORD_MANAGER_NAME. All rights reserved.
								</p>
								</div>
							</td>
							</tr>
						</tbody>
						</table>
					</td>
					</tr>
				</tbody>
				</table>
			</div>
		</div>
	`
}

func ChangePasswordRequest(masterPassword, email string) string {
	return `
	<div>
	<div>
		<table
		width="100%"
		border="0"
		cellspacing="0"
		cellpadding="0"
		style="width: 100% !important"
		>
		<tbody>
			<tr>
			<td align="center">
				<table
				style="
					border: 1px solid #eaeaea;
					border-radius: 5px;
					margin: 40px 0;
				"
				width="600"
				border="0"
				cellspacing="0"
				cellpadding="40"
				>
				<tbody>
					<tr>
					<td align="center">
						<div style="text-align: left; width: 465px">
						<table
							width="100%"
							border="0"
							cellspacing="0"
							cellpadding="0"
							style="width: 100% !important"
						>
							<tbody>
							<tr>
								<td align="center">
								<img
									src="$LOGO_URL"
									alt="$PASSWORD_MANAGER_NAME"
								/>
								<h1
									style="
									color: #000;
									font-family: -apple-system, BlinkMacSystemFont,
										'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
										'Cantarell', 'Fira Sans', 'Droid Sans',
										'Helvetica Neue', sans-serif;
									font-size: 24px;
									font-weight: normal;
									margin: 30px 0;
									padding: 0;
									"
								>
									Change Master Password Request
								</h1>
								</td>
							</tr>
							</tbody>
						</table>

						<p
							style="
							color: #000;
							font-family: -apple-system, BlinkMacSystemFont,
								'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
								'Cantarell', 'Fira Sans', 'Droid Sans',
								'Helvetica Neue', sans-serif;
							font-size: 14px;
							line-height: 24px;
							"
						>
							$PASSWORD_MANAGER_NAME User ` + email + ` want to change his/her master
							password

							<br />
							<br />

							The new master password is
						</p>

						<table
							width="100%"
							border="0"
							cellspacing="0"
							cellpadding="0"
							style="width: 100% !important"
						>
							<tbody>
							<tr>
								<td
								align="center"
								valign="middle"
								style="
									color: #eeeeee !important;
									background-color: #1163e6;
									border-radius: 5px;
									padding: 5px;
									font: 300 20px 'SFMono-Regular', Consolas,
									'Liberation Mono', Menlo, monospace;
								"
								>
								` + masterPassword + `
								</td>
							</tr>
							</tbody>
						</table>

						<br />

						<hr
							style="
							border: none;
							border-top: 1px solid #eaeaea;
							margin: 26px 0;
							width: 100%;
							"
						/>
						<p
							style="
							color: #666666;
							font-family: -apple-system, BlinkMacSystemFont,
								'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
								'Cantarell', 'Fira Sans', 'Droid Sans',
								'Helvetica Neue', sans-serif;
							font-size: 12px;
							line-height: 24px;
							"
						>
							&copy; $PASSWORD_MANAGER_NAME. All rights reserved.
						</p>
						</div>
					</td>
					</tr>
				</tbody>
				</table>
			</td>
			</tr>
		</tbody>
		</table>
	</div>
	</div>`
}
